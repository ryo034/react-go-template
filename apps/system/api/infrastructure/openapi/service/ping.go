package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) APIV1PingGet(ctx context.Context) (openapi.APIV1PingGetRes, error) {
	return &openapi.APIV1PingGetOK{
		Message: openapi.OptString{Value: "pong", Set: true},
	}, nil
}
