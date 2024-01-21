package auth

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/driver/email"
	"github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"github.com/uptrace/bun"
)

type UseCase interface {
	AuthByTOTP(ctx context.Context, input ByTOTPInput) (openapi.OtpAuthPostRes, error)
	VerifyTOTP(ctx context.Context, input VerifyTOTPInput) (openapi.OtpVerifyPostRes, error)
}

type useCase struct {
	txp     core.TransactionProvider
	dbp     core.Provider
	repo    auth.Repository
	emailDr email.Driver
	fbDr    firebase.Driver
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, acRepo auth.Repository, emailDr email.Driver, fbDr firebase.Driver) UseCase {
	return &useCase{txp, dbp, acRepo, emailDr, fbDr}
}

func (u *useCase) AuthByTOTP(ctx context.Context, input ByTOTPInput) (openapi.OtpAuthPostRes, error) {
	code, err := u.repo.GenTOTP(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if err = u.emailDr.Send(ctx, input.Email); err != nil {
		return nil, err
	}
	return &openapi.OtpAuthPostOK{Code: code}, nil
}

func (u *useCase) VerifyTOTP(ctx context.Context, input VerifyTOTPInput) (openapi.OtpVerifyPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (string, error) {
		usr, err := u.repo.Find(pr, p, input.Email)
		if err != nil {
			return "", err
		}
		// if user exists, verify TOTP and return custom token
		if usr != nil {
			return u.verifyTOTP(pr, usr.AccountID(), input.Email, input.Otp)
		}
		// if user does not exist, create user, verify TOTP and return custom token
		return u.verifyTOTPWithCreate(pr, p, input.Email, input.Otp)
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	return &openapi.OtpVerifyPostOK{}, nil
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
	aID := account.GenerateID()
	if _, err := u.repo.Create(ctx, exec, aID, email); err != nil {
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
	return u.fbDr.CustomToken(ctx, aID)
}
