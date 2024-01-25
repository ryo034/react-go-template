package service

import (
	"context"
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
