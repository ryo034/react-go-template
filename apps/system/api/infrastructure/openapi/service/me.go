package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) MeGet(ctx context.Context) (openapi.MeGetRes, error) {
	return &openapi.Me{
		EmailVerified: false,
		MultiFactor: openapi.OptMultiFactor{
			Value: openapi.MultiFactor{},
			Set:   false,
		},
		Member: openapi.Member{
			IdNumber: openapi.OptString{},
			User:     openapi.User{},
		},
	}, nil
}
