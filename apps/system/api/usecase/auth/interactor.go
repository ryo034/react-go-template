package auth

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ryo034/react-go-template/apps/system/api/domain/notification"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"

	"github.com/go-faster/errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"github.com/uptrace/bun"
)

type UseCase interface {
	AuthByOTP(ctx context.Context, i ByOTPInput) (openapi.APIV1AuthOtpPostRes, error)
	AuthByOAuth(ctx context.Context, i ByOAuthInput) (openapi.APIV1AuthOAuthPostRes, error)
	VerifyOTP(ctx context.Context, i VerifyOTPInput) (openapi.APIV1AuthOtpVerifyPostRes, error)
	ProcessInvitationEmail(ctx context.Context, i ProcessInvitationEmailInput) (openapi.ProcessInvitationEmailRes, error)
	ProcessInvitationOAuth(ctx context.Context, i ProcessInvitationOAuthInput) (openapi.ProcessInvitationOAuthRes, error)
	InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.GetInvitationByTokenRes, error)
}

type useCase struct {
	txp              core.TransactionProvider
	dbp              core.Provider
	repo             auth.Repository
	meRepo           me.Repository
	invRepo          invitation.Repository
	wRepo            workspace.Repository
	notificationRepo notification.Repository
	op               OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo auth.Repository, meRepo me.Repository, invRepo invitation.Repository, wRepo workspace.Repository, notificationRepo notification.Repository, op OutputPort) UseCase {
	return &useCase{txp, dbp, acRepo, meRepo, invRepo, wRepo, notificationRepo, op}
}

func (u *useCase) AuthByOTP(ctx context.Context, i ByOTPInput) (openapi.APIV1AuthOtpPostRes, error) {
	code, err := u.repo.GenOTP(ctx, i.Email)
	if err != nil {
		return nil, err
	}
	if err = u.notificationRepo.NotifyOtpByEmail(ctx, i.Email, code); err != nil {
		return nil, err
	}
	return &openapi.APIV1AuthOtpPostOK{}, nil
}

func (u *useCase) createUser(ctx context.Context, p bun.IDB, ci CreateInfo) (*me.Me, error) {
	usr, err := u.repo.Create(ctx, p, ci.User, ci.Provider)
	if err != nil {
		return nil, err
	}
	if err = u.meRepo.UpdateName(ctx, p, usr); err != nil {
		return nil, err
	}
	res, err := u.meRepo.FindBeforeOnboard(ctx, p, usr.AccountID())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *useCase) AuthByOAuth(ctx context.Context, i ByOAuthInput) (openapi.APIV1AuthOAuthPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)

	m, err := u.meRepo.FindLastLogin(ctx, p, i.AccountID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if m == nil {
		if m, err = u.createUser(ctx, p, i.CreateInfo); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	if m.NotJoined() {
		return u.op.AuthByAuth(m)
	}
	if err = u.meRepo.SetMe(ctx, m); err != nil {
		return nil, err
	}
	if err = u.meRepo.RecordLogin(ctx, p, m); err != nil {
		return nil, err
	}
	return u.op.AuthByAuth(m)
}

func (u *useCase) setupUserByOAuth(ctx context.Context, aID account.ID) (context.Context, *provider.Provider, error) {
	ap, err := provider.NewProviderAsEmailOnFirebase(aID)
	if err != nil {
		return ctx, nil, err
	}
	ctx = u.meRepo.SetCurrentProvider(ctx, ap)
	return ctx, ap, nil
}

func (u *useCase) setupOAuthUserByOAuth(ctx context.Context, p bun.IDB, usr *user.User) (context.Context, *me.Me, error) {
	meRes, err := u.meRepo.FindLastLogin(ctx, p, usr.AccountID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx, nil, domainErr.NewUnauthenticated(fmt.Sprintf("User not found: %s", usr.Email().ToString()))
		}
		return ctx, nil, err
	}
	ap := meRes.Providers().FindByKind(provider.Email)
	ctx = u.meRepo.SetCurrentProvider(ctx, ap)
	if err = u.meRepo.SetMe(ctx, meRes); err != nil {
		return ctx, nil, err
	}
	return ctx, meRes, nil
}

// VerifyOTP User information cannot be retrieved/edited in Firebase on the backend side until VerifyOTP returns a token to the frontend and is authenticated
func (u *useCase) VerifyOTP(ctx context.Context, i VerifyOTPInput) (openapi.APIV1AuthOtpVerifyPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	usr, err := u.repo.FindByEmail(ctx, p, i.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	newAccountID, err := account.GenerateID()
	if err != nil {
		return nil, err
	}
	var meRes *me.Me = nil
	var ap *provider.Provider = nil

	//set user info to context before transaction
	if usr == nil {
		ctx, ap, err = u.setupUserByOAuth(ctx, newAccountID)
	} else {
		ctx, meRes, err = u.setupOAuthUserByOAuth(ctx, p, usr)
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}

	fn := func() (string, error) {
		if usr == nil {
			_, err = u.repo.Create(ctx, p, user.NewUser(newAccountID, i.Email, nil, nil, nil), ap)
			if err != nil {
				return "", err
			}
		} else {
			if err = u.meRepo.RecordLogin(ctx, p, meRes); err != nil {
				return "", err
			}
		}
		// if user exists, verify TOTP and return custom token
		return u.repo.VerifyOTP(pr, i.Email, i.Otp)
	}

	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	tk := result.Value(0).(string)
	return u.op.JwtToken(tk), nil
}

func (u *useCase) ProcessInvitationEmail(ctx context.Context, i ProcessInvitationEmailInput) (openapi.ProcessInvitationEmailRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	invRes, err := u.invRepo.FindActiveByEmail(ctx, p, i.Email)
	if err != nil {
		return nil, err
	}
	if err = invRes.ValidateCanVerify(i.Token); err != nil {
		return nil, err
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() error {
		if err = u.invRepo.VerifyByToken(pr, p, i.Token); err != nil {
			return err
		}
		code, err := u.repo.GenOTP(pr, i.Email)
		if err != nil {
			return err
		}
		return u.notificationRepo.NotifyOtpByEmail(pr, i.Email, code)
	}
	return &openapi.ProcessInvitationEmailOK{}, pr.Transactional(fn)().Error()
}

func (u *useCase) ProcessInvitationOAuth(ctx context.Context, i ProcessInvitationOAuthInput) (openapi.ProcessInvitationOAuthRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	invRes, err := u.invRepo.FindActiveByEmail(ctx, p, i.Email)
	if err != nil {
		return nil, err
	}
	if err = invRes.ValidateCanVerify(i.Token); err != nil {
		return nil, err
	}

	// check Account Exists
	usr, err := u.repo.FindByEmail(ctx, p, i.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*me.Me, error) {
		if err = u.invRepo.VerifyByToken(pr, p, i.Token); err != nil {
			return nil, err
		}
		var meRes *me.Me = nil
		if usr != nil {
			if meRes, err = u.meRepo.FindLastLogin(pr, p, usr.AccountID()); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
		}
		if meRes == nil {
			if i.CreateInfo == nil {
				return nil, domainErr.NewUnauthenticated("User not found")
			}
			if meRes, err = u.createUser(pr, p, *i.CreateInfo); err != nil {
				return nil, err
			}
		}
		if meRes.NotJoined() {
			return meRes, nil
		}
		if err = u.meRepo.RecordLogin(pr, p, meRes); err != nil {
			return nil, err
		}
		return meRes, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.ProcessInvitationOAuth(result.Value(0).(*me.Me))
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
