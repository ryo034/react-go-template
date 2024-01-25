package auth

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	authUc "github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
)

type Controller interface {
	AuthByOTP(ctx context.Context, req *openapi.APIV1AuthOtpPostReq) (openapi.APIV1AuthOtpPostRes, error)
	VerifyOTP(ctx context.Context, req *openapi.APIV1AuthOtpVerifyPostReq) (openapi.APIV1AuthOtpVerifyPostRes, error)
	AuthByOAuth(ctx context.Context) (openapi.APIV1AuthOAuthPostRes, error)
}

type controller struct {
	auc  authUc.UseCase
	resl shared.Resolver
}

func NewController(auc authUc.UseCase, resl shared.Resolver) Controller {
	return &controller{auc, resl}
}

func (c *controller) AuthByOTP(ctx context.Context, req *openapi.APIV1AuthOtpPostReq) (openapi.APIV1AuthOtpPostRes, error) {
	em, err := account.NewEmail(req.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOtpPostRes), nil
	}
	inp := authUc.ByTOTPInput{Email: em}
	res, err := c.auc.AuthByTOTP(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOtpPostRes), nil
	}
	return res, nil
}

func (c *controller) VerifyOTP(ctx context.Context, req *openapi.APIV1AuthOtpVerifyPostReq) (openapi.APIV1AuthOtpVerifyPostRes, error) {
	em, err := account.NewEmail(req.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOtpVerifyPostRes), nil
	}

	inp := authUc.VerifyTOTPInput{Email: em, Otp: req.Otp}
	res, err := c.auc.VerifyTOTP(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOtpVerifyPostRes), nil
	}
	return res, nil
}

func (c *controller) AuthByOAuth(ctx context.Context) (openapi.APIV1AuthOAuthPostRes, error) {
	//ctxからtokenを取得
	//em, err := account.NewEmail(req.Email)
	//if err != nil {
	//	return c.resl.Error(ctx, err).(openapi.APIV1AuthOAuthPostRes), nil
	//}
	//
	//inp := authUc.VerifyTOTPInput{Email: em, Otp: req.Otp}
	//res, err := c.auc.VerifyOTP(ctx, inp)
	//if err != nil {
	//	return c.resl.Error(ctx, err).(openapi.APIV1AuthOAuthPostRes), nil
	//}
	return nil, nil
}
