package member

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(m *member.Member) openapi.Member
	AdaptAll(ms member.Members) openapi.Members
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
	dn := ""
	if m.HasDisplayName() {
		dn = m.DisplayName().ToString()
	}
	return openapi.Member{
		Profile: openapi.MemberProfile{
			ID:          m.ID().ToFriendlyString(),
			DisplayName: dn,
			IdNumber: openapi.OptString{
				Value: m.IDNumber().ToString(),
				Set:   m.HasIDNumber(),
			},
		},
		User: a.ua.Adapt(m.User()),
	}
}

func (a *adapter) AdaptAll(ms member.Members) openapi.Members {
	res := make(openapi.Members, ms.Size())
	for i, m := range ms.AsSlice() {
		res[i] = a.Adapt(m)
	}
	return res
}
