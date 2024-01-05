package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) HealthGet(ctx context.Context) (openapi.HealthGetRes, error) {
	return &openapi.HealthGetOK{Status: openapi.HealthGetOKStatusHealthy}, nil
}
