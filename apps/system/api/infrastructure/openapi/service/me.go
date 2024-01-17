package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) MeGet(ctx context.Context) (openapi.MeGetRes, error) {
	aID := ""
	return s.ctrl.Me.Find(ctx, aID)
}
