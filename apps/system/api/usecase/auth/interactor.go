package auth

import (
	"context"
	"database/sql"
	"github.com/go-faster/errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	"github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"github.com/uptrace/bun"
)

type UseCase interface {
	AuthByOTP(ctx context.Context, input ByOTPInput) (openapi.APIV1AuthOtpPostRes, error)
	VerifyOTP(ctx context.Context, input VerifyOTPInput) (openapi.APIV1AuthOtpVerifyPostRes, error)
}

type useCase struct {
	txp     core.TransactionProvider
	dbp     core.Provider
	repo    auth.Repository
	meRepo  me.Repository
	wRepo   workspace.Repository
	emailDr email.Driver
	fbDr    firebase.Driver
	op      OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo auth.Repository, meRepo me.Repository, wRepo workspace.Repository, emailDr email.Driver, fbDr firebase.Driver, op OutputPort) UseCase {
	return &useCase{txp, dbp, acRepo, meRepo, wRepo, emailDr, fbDr, op}
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
		usr, err := u.repo.Find(pr, p, i.Email)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return "", err
		}
		if usr == nil {
			// if user does not exist, create user, verify TOTP and return custom token
			aID, err := account.GenerateID()
			if err != nil {
				return "", err
			}
			usr, err = u.repo.Create(pr, p, aID, i.Email)
			if err != nil {
				return "", err
			}
		}

		// if user exists, verify TOTP and return custom token
		tk, err := u.verifyOTP(pr, usr.AccountID(), i.Email, i.Otp)
		if err != nil {
			return "", err
		}

		//TODO: tkを一度返して、フロント側で認証したらjwtにcurrentWorkspaceIdがセットできるようになる

		// check active invitation
		invRes, wRes, err := u.wRepo.FindActiveInvitation(pr, p, i.Email)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return "", err
		}
		if errors.Is(err, sql.ErrNoRows) || invRes == nil {
			if err = u.memberLastLogin(pr, p, usr.AccountID()); err != nil {
				return "", err
			}
			return tk, nil
		}

		// if the user not joined the workspace, add the user to the workspace
		mem, err := u.wRepo.FindMember(pr, p, usr.AccountID(), wRes.ID())
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return "", err
		}

		if errors.Is(err, sql.ErrNoRows) || mem == nil {
			meRes, err := u.meRepo.FindBeforeOnboard(pr, p, usr.AccountID())
			if err != nil {
				return "", err
			}
			meID, err := member.GenerateID()
			if err != nil {
				return "", err
			}
			n := ""
			if meRes.Self().Name() != nil {
				n = meRes.Self().Name().ToString()
			}
			dn, err := member.NewDisplayName(n)
			if err != nil {
				return "", err
			}
			m := member.NewMember(meID, meRes.Self(), dn, nil)
			mem, err = u.wRepo.AddMember(pr, p, wRes, m)
			if err != nil {
				return "", err
			}
		}

		//meRes, err := u.meRepo.Find(pr, p, mem.ID())
		//if err != nil {
		//	return "", err
		//}
		//if err = u.meRepo.LastLogin(pr, p, meRes); err != nil {
		//	return "", err
		//}
		return tk, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return u.op.JwtToken(result.Value(0).(string)), nil
}

func (u *useCase) verifyOTP(ctx context.Context, aID account.ID, email account.Email, code string) (string, error) {
	ok, err := u.repo.VerifyOTP(ctx, email, code)
	if err != nil {
		return "", err
	}
	//TODO: custom error
	if !ok {
		return "", err
	}
	return u.fbDr.CustomToken(ctx, aID)
}
