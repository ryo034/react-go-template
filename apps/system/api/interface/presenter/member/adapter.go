package member

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/user"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(m *member.Member) openapi.Member
	AdaptAll(ms member.Members) []openapi.Member
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

	p := m.Profile()
	dn := ""
	if p.HasDisplayName() {
		dn = p.DisplayName().ToString()
	}
	return openapi.Member{
		ID:   m.ID().ToFriendlyString(),
		User: a.ua.Adapt(m.User()),
		Profile: openapi.MemberProfile{
			DisplayName: dn,
			Bio:         openapi.OptString{Value: p.Bio().ToString(), Set: p.HasBio()},
			IdNumber: openapi.OptString{
				Value: p.IDNumber().ToString(),
				Set:   p.HasIDNumber(),
			},
		},
	}
}

func (a *adapter) AdaptAll(ms member.Members) []openapi.Member {
	res := make([]openapi.Member, ms.Size())
	for i, m := range ms.AsSlice() {
		res[i] = a.Adapt(m)
	}
	return res
}
