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

func (a *adapter) adaptRole(r member.Role) openapi.MemberRole {
	switch r {
	case member.RoleOwner:
		return openapi.MemberRoleOWNER
	case member.RoleAdmin:
		return openapi.MemberRoleADMIN
	case member.RoleMember:
		return openapi.MemberRoleMEMBER
	default:
		return openapi.MemberRoleGUEST
	}
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

	var memIDNum = openapi.OptString{Set: false}
	if p.HasIDNumber() {
		memIDNum.Set = true
		memIDNum.Value = p.IDNumber().ToString()
	}

	var memBio = openapi.OptString{Set: false}
	if p.HasBio() {
		memBio.Set = true
		memBio.Value = p.Bio().ToString()
	}

	return openapi.Member{
		ID:   m.ID().Value(),
		User: a.ua.Adapt(m.User()),
		Role: a.adaptRole(m.Role()),
		Profile: openapi.MemberProfile{
			DisplayName: dn,
			Bio:         memBio,
			IdNumber:    memIDNum,
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
