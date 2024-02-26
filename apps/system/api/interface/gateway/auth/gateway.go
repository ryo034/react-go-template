package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	authDr "github.com/ryo034/react-go-template/apps/system/api/driver/auth"
	kvDr "github.com/ryo034/react-go-template/apps/system/api/driver/keyvalue"
	auth2 "github.com/ryo034/react-go-template/apps/system/api/infrastructure/auth"
	"github.com/uptrace/bun"
)

type gateway struct {
	kvd kvDr.Store
	ad  authDr.Driver
	adp Adapter
}

func NewGateway(kvd kvDr.Store, ad authDr.Driver, adp Adapter) auth.Repository {
	return &gateway{kvd, ad, adp}
}

const totpKeyPrefix = "otp:"
const totpKeyExpire = 60 * 2

func (g *gateway) GenTOTP(ctx context.Context, email account.Email) (string, error) {
	code, err := auth2.GenerateOTPCode()
	if err != nil {
		return "", err
	}
	if err = g.kvd.Set(ctx, fmt.Sprintf("%s%s", totpKeyPrefix, email.ToString()), code, totpKeyExpire*time.Second); err != nil {
		return "", err
	}
	return code, nil
}

func (g *gateway) VerifyOTP(ctx context.Context, email account.Email, code string) (bool, error) {
	c, err := g.kvd.Get(ctx, fmt.Sprintf("%s%s", totpKeyPrefix, email.ToString()))
	if err != nil {
		return false, err
	}
	return c == code, nil
}

func (g *gateway) Create(ctx context.Context, exec bun.IDB, aID account.ID, email account.Email, ap *provider.Provider) (*user.User, error) {
	res, err := g.ad.Create(ctx, exec, aID, email, ap)
	if err != nil {
		return nil, err
	}
	return g.adp.AdaptTmp(res)
}

func (g *gateway) FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*user.User, error) {
	res, err := g.ad.Find(ctx, exec, email)
	if err != nil {
		return nil, err
	}
	return g.adp.AdaptTmp(res)
}
