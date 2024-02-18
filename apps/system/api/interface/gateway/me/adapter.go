package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	memberGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	userGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
	workspaceGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace"
	invitationGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace/invitation"
)

type Adapter interface {
	Adapt(m *models.Member, ws models.Workspaces, ris []*models.Invitation) (*me.Me, error)
	AdaptSystemAccount(m *models.SystemAccount) (*me.Me, error)
	AdaptReceivedInvitation(i *models.Invitation) (me.ReceivedInvitation, error)
	AdaptAllReceivedInvitation(is []*models.Invitation) (me.ReceivedInvitations, error)
}

type adapter struct {
	uga  userGw.Adapter
	wga  workspaceGw.Adapter
	mga  memberGw.Adapter
	inva invitationGw.Adapter
}

func NewAdapter(uga userGw.Adapter, wga workspaceGw.Adapter, mga memberGw.Adapter, inva invitationGw.Adapter) Adapter {
	return &adapter{uga, wga, mga, inva}
}

func (a *adapter) Adapt(m *models.Member, ws models.Workspaces, ris []*models.Invitation) (*me.Me, error) {
	u, err := a.uga.AdaptTmp(m.SystemAccount)
	mem, err := a.mga.Adapt(m)
	if err != nil {
		return nil, err
	}
	w, err := a.wga.Adapt(m.Workspace)
	if err != nil {
		return nil, err
	}
	aws, err := a.wga.AdaptAll(ws)
	if err != nil {
		return nil, err
	}
	aris, err := a.AdaptAllReceivedInvitation(ris)
	if err != nil {
		return nil, err
	}
	return me.NewMe(u, w, mem, aws, aris), nil
}

func (a *adapter) AdaptSystemAccount(m *models.SystemAccount) (*me.Me, error) {
	u, err := a.uga.AdaptTmp(m)
	if err != nil {
		return nil, err
	}
	return me.NewMe(u, nil, nil, nil, nil), nil
}

func (a *adapter) AdaptReceivedInvitation(i *models.Invitation) (me.ReceivedInvitation, error) {
	inviter, err := a.wga.AdaptInviter(i.InvitationUnit.Workspace, i.InvitationUnit.Member)
	if err != nil {
		return me.ReceivedInvitation{}, err
	}

	inv, err := a.inva.Adapt(i)
	if err != nil {
		return me.ReceivedInvitation{}, err
	}
	return me.NewReceivedInvitation(inv, inviter), nil
}

func (a *adapter) AdaptAllReceivedInvitation(is []*models.Invitation) (me.ReceivedInvitations, error) {
	mis := make([]me.ReceivedInvitation, 0, len(is))
	for _, i := range is {
		ai, err := a.AdaptReceivedInvitation(i)
		if err != nil {
			return nil, err
		}
		mis = append(mis, ai)
	}
	return me.NewReceivedInvitations(mis), nil
}
