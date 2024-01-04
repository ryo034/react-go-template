package openapi

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) HealthGet(ctx context.Context) (openapi.HealthGetRes, error) {
	return &openapi.Health{Status: openapi.HealthStatusHealthy}, nil
}

func (s *service) HogeGet(ctx context.Context) (openapi.HogeGetRes, error) {
	return &openapi.Health{Status: openapi.HealthStatusHealthy}, nil
}
