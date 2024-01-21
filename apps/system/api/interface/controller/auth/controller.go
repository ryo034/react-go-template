package auth

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	authUc "github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
)

type Controller interface {
	AuthByTOTP(ctx context.Context, req *openapi.OtpAuthPostReq) (openapi.OtpAuthPostRes, error)
	VerifyTOTP(ctx context.Context, req *openapi.OtpVerifyPostReq) (openapi.OtpVerifyPostRes, error)
}

type controller struct {
	auc  authUc.UseCase
	resl shared.Resolver
}

func NewController(auc authUc.UseCase, resl shared.Resolver) Controller {
	return &controller{auc, resl}
}

func (c *controller) AuthByTOTP(ctx context.Context, req *openapi.OtpAuthPostReq) (openapi.OtpAuthPostRes, error) {
	em, err := account.NewEmail(req.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.OtpAuthPostRes), nil
	}
	inp := authUc.ByTOTPInput{Email: em}
	res, err := c.auc.AuthByTOTP(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.OtpAuthPostRes), nil
	}
	return res, nil
}

func (c *controller) VerifyTOTP(ctx context.Context, req *openapi.OtpVerifyPostReq) (openapi.OtpVerifyPostRes, error) {
	em, err := account.NewEmail(req.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.OtpVerifyPostRes), nil
	}

	inp := authUc.VerifyTOTPInput{Email: em, Otp: req.Otp}
	res, err := c.auc.VerifyTOTP(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.OtpVerifyPostRes), nil
	}
	return res, nil
}
