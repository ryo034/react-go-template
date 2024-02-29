package me

import (
	"fmt"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	providerDomain "github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	memberGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	userGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
	workspaceGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace"
	invitationGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/workspace/invitation"
)

type Adapter interface {
	Adapt(m *models.Member, ws models.Workspaces, ris []*models.Invitation) (*me.Me, error)
	AdaptSystemAccount(sa *models.SystemAccount) (*me.Me, error)
	AdaptReceivedInvitation(i *models.Invitation) (me.ReceivedInvitation, error)
	AdaptAllReceivedInvitation(is []*models.Invitation) (me.ReceivedInvitations, error)
	AdaptProvider(p *models.AuthProvider) (*providerDomain.Provider, error)
	AdaptAllProviders(ps []*models.AuthProvider) (providerDomain.Providers, error)
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
	prvs, err := a.AdaptAllProviders(m.SystemAccount.AuthProviders)
	if err != nil {
		return nil, err
	}
	return me.NewMe(u, w, mem, aws, aris, prvs), nil
}

func (a *adapter) AdaptSystemAccount(sa *models.SystemAccount) (*me.Me, error) {
	u, err := a.uga.AdaptTmp(sa)
	if err != nil {
		return nil, err
	}
	prvs, err := a.AdaptAllProviders(sa.AuthProviders)
	if err != nil {
		return nil, err
	}
	return me.NewMe(u, nil, nil, nil, nil, prvs), nil
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
	return me.NewReceivedInvitation(inv, inviter)
}

func (a *adapter) AdaptAllReceivedInvitation(is []*models.Invitation) (me.ReceivedInvitations, error) {
	if is == nil {
		return me.NewReceivedInvitations(nil), nil
	}
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

func (a *adapter) AdaptProvider(p *models.AuthProvider) (*providerDomain.Provider, error) {
	prb := providerDomain.ProvidedByFirebase
	var kind = providerDomain.Google
	switch p.Provider {
	case "email":
		kind = providerDomain.Email
	case "google":
		kind = providerDomain.Google
	default:
		return nil, fmt.Errorf("unknown provider: %s", p.Provider)
	}

	apUID, err := providerDomain.NewUID(p.ProviderUID)
	if err != nil {
		return nil, err
	}
	return providerDomain.NewProvider(providerDomain.NewIDFromUUID(p.AuthProviderID), kind, prb, apUID), nil
}

func (a *adapter) AdaptAllProviders(ps []*models.AuthProvider) (providerDomain.Providers, error) {
	res := make([]*providerDomain.Provider, 0, len(ps))
	for _, p := range ps {
		pv, err := a.AdaptProvider(p)
		if err != nil {
			return nil, err
		}
		res = append(res, pv)
	}
	return providerDomain.NewProviders(res), nil
}
