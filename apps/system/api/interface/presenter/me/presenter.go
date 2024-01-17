package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

func NewPresenter() meUc.OutputPort {
	return &presenter{}
}

type presenter struct{}

func (p *presenter) Find(m *me.Me) *openapi.Me {
	return &openapi.Me{}
}
