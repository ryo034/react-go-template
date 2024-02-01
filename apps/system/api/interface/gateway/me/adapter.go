package me

import (
	"firebase.google.com/go/v4/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	memberGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	userGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
	workspaceGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace"
)

type Adapter interface {
	Adapt(m *models.Member, fu *auth.UserRecord) (*me.Me, error)
	AdaptSystemAccount(m *models.SystemAccount) (*me.Me, error)
}

type adapter struct {
	uga userGw.Adapter
	wga workspaceGw.Adapter
	mga memberGw.Adapter
}

func NewAdapter(uga userGw.Adapter, wga workspaceGw.Adapter, mga memberGw.Adapter) Adapter {
	return &adapter{uga, wga, mga}
}

func (a *adapter) Adapt(m *models.Member, fu *auth.UserRecord) (*me.Me, error) {
	u, err := a.uga.Adapt(m.SystemAccount)
	mem, err := a.mga.Adapt(m)
	if err != nil {
		return nil, err
	}
	w, err := a.wga.Adapt(m.Workspace)
	if err != nil {
		return nil, err
	}
	return me.NewMe(u, w, mem), nil
}

func (a *adapter) AdaptSystemAccount(m *models.SystemAccount) (*me.Me, error) {
	u, err := a.uga.Adapt(m)
	if err != nil {
		return nil, err
	}
	return me.NewMe(u, nil, nil), nil
}
