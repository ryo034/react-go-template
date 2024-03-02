package auth

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/driver/firebase"

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
	ProcessInvitationEmail(ctx context.Context, i ProcessInvitationInput) (openapi.ProcessInvitationEmailRes, error)
	ProcessInvitationOAuth(ctx context.Context, i ProcessInvitationOAuth) (openapi.ProcessInvitationOAuthRes, error)
	InvitationByToken(ctx context.Context, i InvitationByTokenInput) (openapi.GetInvitationByTokenRes, error)
}

type controller struct {
	auc  authUc.UseCase
	resl shared.Resolver
	co   infraShared.ContextOperator
	fbDr firebase.Driver
}

type ProcessInvitationInput struct {
	Token uuid.UUID
	Email string
}

type ProcessInvitationOAuth struct {
	Token uuid.UUID
}

type InvitationByTokenInput struct {
	Token uuid.UUID
}

func NewController(auc authUc.UseCase, resl shared.Resolver, co infraShared.ContextOperator, fbDr firebase.Driver) Controller {
	return &controller{auc, resl, co, fbDr}
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
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOAuthPostRes), nil
	}
	apUID, err := c.co.GetAuthProviderUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOAuthPostRes), nil
	}
	inp := authUc.ByOAuthInput{AccountID: aID, AuthProviderUID: apUID}
	res, err := c.auc.AuthByOAuth(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1AuthOAuthPostRes), nil
	}
	return res, nil
}

func (c *controller) ProcessInvitationEmail(ctx context.Context, i ProcessInvitationInput) (openapi.ProcessInvitationEmailRes, error) {
	em, err := account.NewEmail(i.Email)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.ProcessInvitationEmailRes), nil
	}
	inp := authUc.ProcessInvitationEmailInput{Token: invitation.NewToken(i.Token), Email: em}
	res, err := c.auc.ProcessInvitationEmail(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.ProcessInvitationEmailRes), nil
	}
	return res, nil
}

func (c *controller) ProcessInvitationOAuth(ctx context.Context, i ProcessInvitationOAuth) (openapi.ProcessInvitationOAuthRes, error) {
	fbUsr, err := c.fbDr.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	var aID *account.ID = nil
	if fbUsr.CustomClaims["account_id"] != nil {
		tmpAID := account.NewIDFromUUID(uuid.MustParse(fbUsr.CustomClaims["account_id"].(string)))
		aID = &tmpAID
	}
	em, _ := account.NewEmail(fbUsr.Email)
	inp := authUc.ProcessInvitationOAuthInput{
		Token:     invitation.NewToken(i.Token),
		Email:     em,
		AccountID: aID,
	}
	res, err := c.auc.ProcessInvitationOAuth(ctx, inp)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.ProcessInvitationOAuthRes), nil
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
