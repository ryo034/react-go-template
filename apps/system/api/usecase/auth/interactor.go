package auth

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"

	"github.com/go-faster/errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	"github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"github.com/uptrace/bun"
)

type UseCase interface {
	AuthByOTP(ctx context.Context, i ByOTPInput) (openapi.APIV1AuthOtpPostRes, error)
	AuthByOAuth(ctx context.Context, i ByOAuthInput) (openapi.APIV1AuthOAuthPostRes, error)
	VerifyOTP(ctx context.Context, i VerifyOTPInput) (openapi.APIV1AuthOtpVerifyPostRes, error)
	ProcessInvitation(ctx context.Context, i ProcessInvitationInput) (openapi.ProcessInvitationRes, error)
	InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.GetInvitationByTokenRes, error)
}

type useCase struct {
	txp     core.TransactionProvider
	dbp     core.Provider
	repo    auth.Repository
	meRepo  me.Repository
	invRepo invitation.Repository
	wRepo   workspace.Repository
	emailDr email.Driver
	fbDr    firebase.Driver
	op      OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo auth.Repository, meRepo me.Repository, invRepo invitation.Repository, wRepo workspace.Repository, emailDr email.Driver, fbDr firebase.Driver, op OutputPort) UseCase {
	return &useCase{txp, dbp, acRepo, meRepo, invRepo, wRepo, emailDr, fbDr, op}
}

func (u *useCase) AuthByOTP(ctx context.Context, i ByOTPInput) (openapi.APIV1AuthOtpPostRes, error) {
	code, err := u.repo.GenTOTP(ctx, i.Email)
	if err != nil {
		return nil, err
	}
	if err = u.emailDr.SendOTP(ctx, i.Email, code); err != nil {
		return nil, err
	}
	return &openapi.APIV1AuthOtpPostOK{}, nil
}

func (u *useCase) createUser(ctx context.Context, p bun.IDB, i ByOAuthInput) (openapi.APIV1AuthOAuthPostRes, error) {
	prov, err := u.fbDr.FindProviderData(ctx)
	if err != nil {
		return nil, err
	}
	fbur, err := u.fbDr.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	ema, err := account.NewEmail(fbur.Email)
	if err != nil {
		return nil, err
	}
	var na *account.Name = nil
	name, err := account.NewName(fbur.DisplayName)
	if err != nil {
		// if not match name format, just ignore
		fmt.Printf("failed to create name: %s", err)
	} else {
		na = &name
	}
	var pn *phone.Number = nil
	tmpPn, err := phone.NewInternationalPhoneNumber(fbur.PhoneNumber, "")
	if err != nil {
		// if not match phone number format, just ignore
		fmt.Printf("failed to create phone number: %s", err)
	} else {
		pn = &tmpPn
	}
	usr, err := u.repo.Create(ctx, p, user.NewUser(i.AccountID, ema, na, pn), prov)
	if err != nil {
		return nil, err
	}
	res, err := u.meRepo.FindBeforeOnboard(ctx, p, usr.AccountID())
	if err != nil {
		return nil, err
	}
	return u.op.AuthByAuth(res)
}

func (u *useCase) AuthByOAuth(ctx context.Context, i ByOAuthInput) (openapi.APIV1AuthOAuthPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)

	//　Account Exists
	m, err := u.meRepo.FindBeforeOnboard(ctx, p, i.AccountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.createUser(ctx, p, i)
		}
		return nil, err
	}

	if m != nil {
		// if user exists and joined return me with member info
		m, err = u.meRepo.FindLastLogin(ctx, p, m.Self().AccountID())
		if err != nil {
			return nil, err
		}

		if m.HasMember() {
			// create last login info and
			if err = u.meRepo.LastLogin(ctx, p, m); err != nil {
				return nil, err
			}
		}
	}
	return u.op.AuthByAuth(m)
}

// memberLastLogin If there is information about the last logged-in workspace,put that workspace information in the jwt.
func (u *useCase) memberLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) error {
	meRes, err := u.meRepo.FindLastLogin(ctx, exec, aID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if errors.Is(err, sql.ErrNoRows) || meRes.NotJoined() {
		return nil
	}
	return u.meRepo.LastLogin(ctx, exec, meRes)
}

func (u *useCase) VerifyOTP(ctx context.Context, i VerifyOTPInput) (openapi.APIV1AuthOtpVerifyPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (string, error) {
		usr, err := u.repo.FindByEmail(pr, p, i.Email)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return "", err
		}
		if usr == nil {
			// if user does not exist, create user, verify TOTP and return custom token
			aID, err := account.GenerateID()
			if err != nil {
				return "", err
			}
			apUID, err := provider.NewUID(aID.Value().String())
			if err != nil {
				return "", err
			}
			ap, err := provider.NewProviderAsEmailOnFirebase(apUID)
			if err != nil {
				return "", err
			}
			usr, err = u.repo.Create(pr, p, user.NewUser(aID, i.Email, nil, nil), ap)
			if err != nil {
				return "", err
			}
		}

		// if user exists, verify TOTP and return custom token
		tk, err := u.verifyOTP(pr, i.Email, i.Otp)
		if err != nil {
			return "", err
		}
		if err = u.memberLastLogin(pr, p, usr.AccountID()); err != nil {
			return "", err
		}
		if tk == "" {
			return "", domainErr.NewUnauthenticated(fmt.Sprintf("Invalid OTP: %s", i.Otp))
		}
		return tk, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	tk := result.Value(0).(string)
	return u.op.JwtToken(tk), nil
}

func (u *useCase) verifyOTP(ctx context.Context, email account.Email, code string) (string, error) {
	ok, err := u.repo.VerifyOTP(ctx, email, code)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", domainErr.NewUnauthenticated(fmt.Sprintf("Invalid OTP: %s", code))
	}
	return u.fbDr.CustomToken(ctx)
}

func (u *useCase) ProcessInvitation(ctx context.Context, i ProcessInvitationInput) (openapi.ProcessInvitationRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	invRes, err := u.invRepo.FindActiveByEmail(ctx, p, i.Email)
	if err != nil {
		return nil, err
	}
	if invRes.Token().NotEquals(i.Token) {
		return nil, invitation.NewInvalidInviteToken(i.Token.Value())
	}
	if err = invRes.ValidateCanVerify(); err != nil {
		return nil, invitation.NewAlreadyExpiredInvitation(invRes.ID(), invRes.Token().Value())
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() error {
		if err = u.invRepo.VerifyByToken(pr, p, i.Token); err != nil {
			return err
		}
		code, err := u.repo.GenTOTP(pr, i.Email)
		if err != nil {
			return err
		}
		return u.emailDr.SendOTP(pr, i.Email, code)
	}
	return &openapi.ProcessInvitationOK{}, pr.Transactional(fn)().Error()
}

func (u *useCase) InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.GetInvitationByTokenRes, error) {
	p := u.dbp.GetExecutor(ctx, true)
	res, err := u.invRepo.FindByToken(ctx, p, i.Token)
	if err != nil {
		return nil, err
	}
	if err = res.ValidateCanGetByToken(); err != nil {
		return nil, err
	}
	w, err := u.wRepo.FindInviterFromToken(ctx, p, i.Token)
	if err != nil {
		return nil, err
	}
	invRes, err := u.invRepo.FindByToken(ctx, p, i.Token)
	if err != nil {
		return nil, err
	}
	ri, err := me.NewReceivedInvitation(invRes, w)
	if err != nil {
		return nil, err
	}
	return u.op.InvitationByToken(ri)
}
