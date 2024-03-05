package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	authDr "github.com/ryo034/react-go-template/apps/system/api/driver/auth"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	kvDr "github.com/ryo034/react-go-template/apps/system/api/driver/keyvalue"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/authentication"
	userGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
	"github.com/uptrace/bun"
)

type gateway struct {
	kvd kvDr.Store
	ad  authDr.Driver
	fd  fbDr.Driver
	ua  userGw.Adapter
}

func NewGateway(kvd kvDr.Store, ad authDr.Driver, fd fbDr.Driver, ua userGw.Adapter) auth.Repository {
	return &gateway{kvd, ad, fd, ua}
}

const otpKeyPrefix = "otp:"
const otpKeyExpire = 60 * 2

func (g *gateway) GenOTP(ctx context.Context, email account.Email) (string, error) {
	code, err := authentication.GenerateOTPCode()
	if err != nil {
		return "", err
	}
	if err = g.kvd.Set(ctx, fmt.Sprintf("%s%s", otpKeyPrefix, email.ToString()), code, otpKeyExpire*time.Second); err != nil {
		return "", err
	}
	return code, nil
}

func (g *gateway) VerifyOTP(ctx context.Context, email account.Email, code string) (string, error) {
	c, err := g.kvd.Get(ctx, fmt.Sprintf("%s%s", otpKeyPrefix, email.ToString()))
	if err != nil {
		return "", err
	}
	if c != code {
		return "", domainErr.NewUnauthenticated(fmt.Sprintf("Invalid OTP: %s", code))
	}
	tk, err := g.fd.CustomToken(ctx)
	if err != nil {
		return "", err
	}
	if tk == "" {
		return "", domainErr.NewUnauthenticated("Failed to generate custom token")
	}
	return tk, nil
}

func (g *gateway) Create(ctx context.Context, exec bun.IDB, usr *user.User, ap *provider.Provider) (*user.User, error) {
	res, err := g.ad.Create(ctx, exec, usr, ap)
	if err != nil {
		return nil, err
	}
	return g.ua.AdaptTmp(res)
}

func (g *gateway) FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*user.User, error) {
	res, err := g.ad.Find(ctx, exec, email)
	if err != nil {
		return nil, err
	}
	return g.ua.AdaptTmp(res)
}
