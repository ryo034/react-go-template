package member

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(m *member.Member) openapi.Member
}

type adapter struct {
	ua user.Adapter
}

func NewAdapter(ua user.Adapter) Adapter {
	return &adapter{ua}
}

func (a *adapter) Adapt(m *member.Member) openapi.Member {
	if m == nil {
		return openapi.Member{}
	}
	return openapi.Member{
		Profile: openapi.MemberProfile{
			ID:          m.ID().ToFriendlyString(),
			DisplayName: m.DisplayName().ToString(),
			IdNumber: openapi.OptString{
				Value: m.IDNumber().ToString(),
				Set:   m.HasIDNumber(),
			},
		},
		User: a.ua.Adapt(m.User()),
	}
}
