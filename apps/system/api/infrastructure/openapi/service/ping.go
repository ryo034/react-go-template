package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) PingGet(ctx context.Context) (openapi.PingGetRes, error) {
	return &openapi.PingGetOK{}, nil
}
