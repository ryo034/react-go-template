package service

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/auth"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1AuthByOtp(ctx context.Context, req *openapi.APIV1AuthByOtpReq) (openapi.APIV1AuthByOtpRes, error) {
	return s.ctrl.Auth.AuthByOTP(ctx, req)
}

func (s *service) APIV1VerifyOTP(ctx context.Context, req *openapi.APIV1VerifyOTPReq) (openapi.APIV1VerifyOTPRes, error) {
	return s.ctrl.Auth.VerifyOTP(ctx, req)
}

func (s *service) APIV1AuthByOAuth(ctx context.Context) (openapi.APIV1AuthByOAuthRes, error) {
	return s.ctrl.Auth.AuthByOAuth(ctx)
}

func (s *service) APIV1ProcessInvitationEmail(ctx context.Context, req *openapi.APIV1ProcessInvitationEmailReq) (openapi.APIV1ProcessInvitationEmailRes, error) {
	return s.ctrl.Auth.ProcessInvitationEmail(ctx, auth.ProcessInvitationInput{
		Token: req.Token,
		Email: req.Email,
	})
}

func (s *service) APIV1ProcessInvitationOAuth(ctx context.Context, req *openapi.APIV1ProcessInvitationOAuthReq) (openapi.APIV1ProcessInvitationOAuthRes, error) {
	return s.ctrl.Auth.ProcessInvitationOAuth(ctx, auth.ProcessInvitationOAuth{
		Token: req.Token,
	})
}

func (s *service) APIV1GetInvitationByToken(ctx context.Context, params openapi.APIV1GetInvitationByTokenParams) (openapi.APIV1GetInvitationByTokenRes, error) {
	return s.ctrl.Auth.InvitationByToken(ctx, auth.InvitationByTokenInput{
		Token: params.Token,
	})
}
