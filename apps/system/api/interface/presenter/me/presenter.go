package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

func NewPresenter(ua user.Adapter, ma member.Adapter, wa workspace.Adapter) meUc.OutputPort {
	return &presenter{ua, ma, wa}
}

type presenter struct {
	ua user.Adapter
	ma member.Adapter
	wa workspace.Adapter
}

func (p *presenter) Find(m *me.Me) *openapi.Me {
	return &openapi.Me{
		Self: p.ua.Adapt(m.Self()),
		Member: openapi.OptMember{
			Set:   m.HasMember(),
			Value: p.ma.Adapt(m.Member()),
		},
		CurrentWorkspace: openapi.OptWorkspace{
			Set:   m.HasWorkspace(),
			Value: p.wa.Adapt(m.Workspace()),
		},
	}
}
