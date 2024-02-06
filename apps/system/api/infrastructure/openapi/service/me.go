package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1MeGet(ctx context.Context) (openapi.APIV1MeGetRes, error) {
	return s.ctrl.Me.Find(ctx)
}

func (s *service) APIV1MeProfilePut(ctx context.Context, req *openapi.APIV1MeProfilePutReq) (openapi.APIV1MeProfilePutRes, error) {
	return s.ctrl.Me.UpdateProfile(ctx, req.User)
}
