package openapi

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/injector"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type service struct {
	openapi.UnimplementedHandler // automatically implement all methods
	inj                          *injector.Injector
}

func NewService(inj *injector.Injector) openapi.Handler {
	return &service{
		openapi.UnimplementedHandler{},
		inj,
	}
}
