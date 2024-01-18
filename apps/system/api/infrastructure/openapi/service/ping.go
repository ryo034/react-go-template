package service

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

func (s *service) PingGet(ctx context.Context) (openapi.PingGetRes, error) {
	//return &openapi.PingGetOK{}, nil
	return &openapi.InternalServerError{
		Status: openapi.OptInt{
			Value: 500,
			Set:   true,
		},
	}, nil
}

func (s *service) PingPost(ctx context.Context, req openapi.OptPingPostReq) (openapi.PingPostRes, error) {
	//return &openapi.PingPostOK{}, nil
	return &openapi.PingPostOK{}, nil
}
