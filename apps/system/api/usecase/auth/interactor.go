package auth

import (
	"context"
	"database/sql"
	"github.com/go-faster/errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	"github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"github.com/uptrace/bun"
)

type UseCase interface {
	AuthByTOTP(ctx context.Context, input ByTOTPInput) (openapi.APIV1AuthOtpPostRes, error)
	VerifyTOTP(ctx context.Context, input VerifyTOTPInput) (openapi.APIV1AuthOtpVerifyPostRes, error)
}

type useCase struct {
	txp     core.TransactionProvider
	dbp     core.Provider
	repo    auth.Repository
	meRepo  me.Repository
	emailDr email.Driver
	fbDr    firebase.Driver
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo auth.Repository, meRepo me.Repository, emailDr email.Driver, fbDr firebase.Driver) UseCase {
	return &useCase{txp, dbp, acRepo, meRepo, emailDr, fbDr}
}

func (u *useCase) AuthByTOTP(ctx context.Context, i ByTOTPInput) (openapi.APIV1AuthOtpPostRes, error) {
	code, err := u.repo.GenTOTP(ctx, i.Email)
	if err != nil {
		return nil, err
	}
	if err = u.emailDr.Send(ctx, i.Email); err != nil {
		return nil, err
	}
	return &openapi.APIV1AuthOtpPostOK{Code: code}, nil
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

func (u *useCase) VerifyTOTP(ctx context.Context, i VerifyTOTPInput) (openapi.APIV1AuthOtpVerifyPostRes, error) {
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
			return u.verifyTOTPWithCreate(pr, p, i.Email, i.Otp)
		}

		// if user exists, verify TOTP and return custom token
		tk, err := u.verifyTOTP(pr, usr.AccountID(), i.Email, i.Otp)
		if err != nil {
			return "", err
		}

		if err = u.memberLastLogin(pr, p, usr.AccountID()); err != nil {
			return "", err
		}
		return tk, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return &openapi.APIV1AuthOtpVerifyPostOK{
		Token: result.Value(0).(string),
	}, nil
}

func (u *useCase) verifyTOTP(ctx context.Context, aID account.ID, email account.Email, code string) (string, error) {
	ok, err := u.repo.VerifyTOTP(ctx, email, code)
	if err != nil {
		return "", err
	}
	//TODO: custom error
	if !ok {
		return "", err
	}
	return u.fbDr.CustomToken(ctx, aID)
}

func (u *useCase) verifyTOTPWithCreate(ctx context.Context, exec bun.IDB, email account.Email, code string) (string, error) {
	aID, err := account.GenerateID()
	if err != nil {
		return "", err
	}
	if _, err = u.repo.Create(ctx, exec, aID, email); err != nil {
		return "", err
	}
	ok, err := u.repo.VerifyTOTP(ctx, email, code)
	if err != nil {
		return "", err
	}
	//TODO: custom error
	if !ok {
		return "", err
	}
	tk, err := u.fbDr.CustomToken(ctx, aID)
	if err != nil {
		return "", err
	}
	if err = u.fbDr.CreateUser(ctx, aID, email); err != nil {
		return "", err
	}
	return tk, nil
}
