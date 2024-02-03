package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/interface/controller/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1MeGet(ctx context.Context) (openapi.APIV1MeGetRes, error) {
	return s.ctrl.Me.Find(ctx)
}

func (s *service) UpdateName(ctx context.Context, req *openapi.UpdateNameReq) (openapi.UpdateNameRes, error) {
	return s.ctrl.Me.UpdateName(ctx, me.UpdateNameInput{
		Name: req.Name,
	})
}
