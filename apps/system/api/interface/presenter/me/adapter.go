package me

import (
	"fmt"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	providerDomain "github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(me *me.Me) (openapi.Me, error)
	AdaptReceivedInvitation(ri me.ReceivedInvitation) (openapi.ReceivedInvitation, error)
	AdaptAllReceivedInvitation(ris me.ReceivedInvitations) ([]openapi.ReceivedInvitation, error)
	AdaptProvider(p *providerDomain.Provider) (openapi.AuthProvider, error)
	AdaptAllProvider(ps providerDomain.Providers) ([]openapi.AuthProvider, error)
}

type adapter struct {
	ua user.Adapter
	ma member.Adapter
	wa workspace.Adapter
	ia invitation.Adapter
}

func NewAdapter(
	ua user.Adapter,
	ma member.Adapter,
	wa workspace.Adapter,
	ia invitation.Adapter,
) Adapter {
	return &adapter{ua, ma, wa, ia}
}

func (a *adapter) Adapt(m *me.Me) (openapi.Me, error) {
	var mem = openapi.OptMember{Set: false}
	if m.IsJoined() {
		mem.Set = true
		mem.Value = a.ma.Adapt(m.Member())
	}

	var cw = openapi.OptWorkspace{Set: false}
	if m.IsJoined() {
		cw.Set = true
		cw.Value = a.wa.Adapt(m.Workspace())
	}

	ris, err := a.AdaptAllReceivedInvitation(m.ReceivedInvitations())
	if err != nil {
		return openapi.Me{}, nil
	}

	providers, err := a.AdaptAllProvider(m.Providers())
	if err != nil {
		return openapi.Me{}, nil
	}
	return openapi.Me{
		Self:                a.ua.Adapt(m.Self()),
		Member:              mem,
		CurrentWorkspace:    cw,
		JoinedWorkspaces:    a.wa.AdaptAll(m.JoinedWorkspaces()),
		ReceivedInvitations: ris,
		Providers:           providers,
	}, nil
}

func (a *adapter) AdaptReceivedInvitation(ri me.ReceivedInvitation) (openapi.ReceivedInvitation, error) {
	i, err := a.ia.Adapt(ri.Invitation())
	if err != nil {
		return openapi.ReceivedInvitation{}, err
	}
	return openapi.ReceivedInvitation{
		Invitation: i,
		Inviter:    a.wa.AdaptInviter(ri.Inviter()),
	}, nil
}

func (a *adapter) AdaptAllReceivedInvitation(ris me.ReceivedInvitations) ([]openapi.ReceivedInvitation, error) {
	if ris == nil || ris.IsEmpty() {
		return nil, nil
	}
	res := make([]openapi.ReceivedInvitation, 0, ris.Size())
	for _, ri := range ris.AsSlice() {
		a, err := a.AdaptReceivedInvitation(ri)
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}

func (a *adapter) AdaptProvider(p *providerDomain.Provider) (openapi.AuthProvider, error) {
	switch p.Kind() {
	case providerDomain.Email:
		return openapi.AuthProviderEmail, nil
	case providerDomain.Google:
		return openapi.AuthProviderGoogle, nil
	}
	return "", fmt.Errorf("unknown provider: %s", p.Kind())
}

func (a *adapter) AdaptAllProvider(ps providerDomain.Providers) ([]openapi.AuthProvider, error) {
	res := make([]openapi.AuthProvider, 0)
	for _, p := range ps.AsSlice() {
		ap, err := a.AdaptProvider(p)
		if err != nil {
			return nil, err
		}
		res = append(res, ap)
	}
	return res, nil
}
