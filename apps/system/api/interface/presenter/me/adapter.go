package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	AdaptReceivedInvitation(ri me.ReceivedInvitation) (openapi.ReceivedInvitation, error)
	AdaptAllReceivedInvitation(ris me.ReceivedInvitations) ([]openapi.ReceivedInvitation, error)
}

type adapter struct {
	wa workspace.Adapter
	ia invitation.Adapter
}

func NewAdapter(wa workspace.Adapter, ia invitation.Adapter) Adapter {
	return &adapter{wa, ia}
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
