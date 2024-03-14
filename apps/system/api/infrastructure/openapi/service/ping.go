package service

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) Ping(ctx context.Context) (openapi.PingRes, error) {
	return &openapi.PingOK{
		Message: openapi.OptString{Value: "pong", Set: true},
	}, nil
}
