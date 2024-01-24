package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1OtpAuthPost(ctx context.Context, req *openapi.APIV1OtpAuthPostReq) (openapi.APIV1OtpAuthPostRes, error) {
	return s.ctrl.Auth.AuthByTOTP(ctx, req)
}

func (s *service) APIV1OtpVerifyPost(ctx context.Context, req *openapi.APIV1OtpVerifyPostReq) (openapi.APIV1OtpVerifyPostRes, error) {
	return s.ctrl.Auth.VerifyTOTP(ctx, req)
}
