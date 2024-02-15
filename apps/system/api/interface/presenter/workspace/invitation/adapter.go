package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(i *invitation.Invitation) (openapi.Invitation, error)
	AdaptAll(is invitation.Invitations) ([]openapi.Invitation, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(i *invitation.Invitation) (openapi.Invitation, error) {
	if i == nil {
		return openapi.Invitation{}, nil
	}
	lt, err := i.ExpiredAt().Value().ToLocalTime()
	if err != nil {
		return openapi.Invitation{}, err
	}
	return openapi.Invitation{
		ID:           i.ID().Value(),
		Verified:     i.Verified(),
		ExpiredAt:    lt,
		InviteeEmail: i.InviteeEmail().ToString(),
		DisplayName:  i.DisplayName().ToString(),
	}, nil
}

func (a *adapter) AdaptAll(is invitation.Invitations) ([]openapi.Invitation, error) {
	if is == nil {
		return nil, nil
	}
	res := make([]openapi.Invitation, 0, is.Size())
	for _, w := range is.AsSlice() {
		a, err := a.Adapt(w)
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}
