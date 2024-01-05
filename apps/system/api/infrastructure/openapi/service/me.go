package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) MeGet(ctx context.Context) (openapi.MeGetRes, error) {
	return &openapi.MeGetOK{
		EmailVerified: false,
		MultiFactor: openapi.OptMeGetOKMultiFactor{
			Value: openapi.MeGetOKMultiFactor{},
			Set:   false,
		},
		User: openapi.MeGetOKUser{
			UserID:      "",
			Email:       "",
			FirstName:   "",
			LastName:    "",
			PhoneNumber: openapi.OptString{},
		},
	}, nil
}
