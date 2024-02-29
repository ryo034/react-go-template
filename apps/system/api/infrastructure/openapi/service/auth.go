package service

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/auth"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1AuthOtpPost(ctx context.Context, req *openapi.APIV1AuthOtpPostReq) (openapi.APIV1AuthOtpPostRes, error) {
	return s.ctrl.Auth.AuthByOTP(ctx, req)
}

func (s *service) APIV1AuthOtpVerifyPost(ctx context.Context, req *openapi.APIV1AuthOtpVerifyPostReq) (openapi.APIV1AuthOtpVerifyPostRes, error) {
	return s.ctrl.Auth.VerifyOTP(ctx, req)
}

func (s *service) APIV1AuthOAuthPost(ctx context.Context) (openapi.APIV1AuthOAuthPostRes, error) {
	return s.ctrl.Auth.AuthByOAuth(ctx)
}

func (s *service) ProcessInvitationEmail(ctx context.Context, req *openapi.ProcessInvitationEmailReq) (openapi.ProcessInvitationEmailRes, error) {
	return s.ctrl.Auth.ProcessInvitationEmail(ctx, auth.ProcessInvitationInput{
		Token: req.Token,
		Email: req.Email,
	})
}

func (s *service) ProcessInvitationOAuth(ctx context.Context, req *openapi.ProcessInvitationOAuthReq) (openapi.ProcessInvitationOAuthRes, error) {
	return s.ctrl.Auth.ProcessInvitationOAuth(ctx, auth.ProcessInvitationOAuth{
		Token: req.Token,
	})
}

func (s *service) GetInvitationByToken(ctx context.Context, params openapi.GetInvitationByTokenParams) (openapi.GetInvitationByTokenRes, error) {
	return s.ctrl.Auth.InvitationByToken(ctx, auth.InvitationByTokenInput{
		Token: params.Token,
	})
}
