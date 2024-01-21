package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) OtpAuthPost(ctx context.Context, req *openapi.OtpAuthPostReq) (openapi.OtpAuthPostRes, error) {
	return s.ctrl.Auth.AuthByTOTP(ctx, req)
}

func (s *service) OtpVerifyPost(ctx context.Context, req *openapi.OtpVerifyPostReq) (openapi.OtpVerifyPostRes, error) {
	return s.ctrl.Auth.VerifyTOTP(ctx, req)
}
