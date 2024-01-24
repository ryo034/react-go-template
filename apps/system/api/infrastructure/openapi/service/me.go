package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1MeGet(ctx context.Context) (openapi.APIV1MeGetRes, error) {
	aID := ""
	wID := ""
	return s.ctrl.Me.Find(ctx, aID, wID)
}
