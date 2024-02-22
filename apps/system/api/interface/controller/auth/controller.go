package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	authUc "github.com/ryo034/react-go-template/apps/system/api/usecase/auth"
)

type Controller interface {
	AuthByOTP(ctx context.Context, req *openapi.APIV1AuthOtpPostReq) (openapi.APIV1AuthOtpPostRes, error)
	VerifyOTP(ctx context.Context, req *openapi.APIV1AuthOtpVerifyPostReq) (openapi.APIV1AuthOtpVerifyPostRes, error)
	AuthByOAuth(ctx context.Context) (openapi.APIV1AuthOAuthPostRes, error)
	ProcessInvitation(ctx context.Context, i ProcessInvitationInput) (openapi.ProcessInvitationRes, error)
	InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.GetInvitationByTokenRes, error)
}

type controller struct {
	auc  authUc.UseCase
	resl shared.Resolver
	co   infraShared.ContextOperator
}

type ProcessInvitationInput struct {
	Token uuid.UUID
	Email string
}

type InvitationByTokenInput struct {
	Token uuid.UUID
}

func NewController(auc authUc.UseCase, resl shared.Resolver, co infraShared.ContextOperator) Controller {
	return &controller{auc, resl, co}
}

func (c *controller) AuthByOTP(ctx context.Context, req *openapi.APIV1AuthOtpPostReq) (openapi.APIV1AuthOtpPostRes, error) {
	em, err := account.NewEmail(req.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOtpPostRes), nil
	}
	inp := authUc.ByOTPInput{Email: em}
	res, err := c.auc.AuthByOTP(ctx, inp)
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
	inp := authUc.VerifyOTPInput{Email: em, Otp: req.Otp}
	return c.auc.VerifyOTP(ctx, inp)
}

func (c *controller) AuthByOAuth(ctx context.Context) (openapi.APIV1AuthOAuthPostRes, error) {
	//ctxからtokenを取得
	//em, err := account.NewEmail(req.Email)
	//if err != nil {
	//	return c.resl.Error(ctx, err).(openapi.APIV1AuthOAuthPostRes), nil
	//}
	//
	//inp := authUc.VerifyOTPInput{Email: em, Otp: req.Otp}
	//res, err := c.auc.VerifyOTP(ctx, inp)
	//if err != nil {
	//	return c.resl.Error(ctx, err).(openapi.APIV1AuthOAuthPostRes), nil
	//}
	return nil, nil
}

func (c *controller) ProcessInvitation(ctx context.Context, i ProcessInvitationInput) (openapi.ProcessInvitationRes, error) {
	em, err := account.NewEmail(i.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.ProcessInvitationRes), nil
	}
	inp := authUc.ProcessInvitationInput{Token: invitation.NewToken(i.Token), Email: em}
	res, err := c.auc.ProcessInvitation(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.ProcessInvitationRes), nil
	}
	return res, nil
}

func (c *controller) InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.GetInvitationByTokenRes, error) {
	inp := authUc.InvitationByTokenInput{Token: invitation.NewToken(i.Token)}
	res, err := c.auc.InvitationByToken(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.GetInvitationByTokenRes), nil
	}
	return res, nil
}
