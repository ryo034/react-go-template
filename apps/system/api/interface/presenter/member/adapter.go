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

	var memStatus = openapi.MemberMembershipStatusLEFT
	if m.MembershipStatus().IsActive() {
		memStatus = openapi.MemberMembershipStatusACTIVE
	}

	p := m.Profile()

	var memIDNum = openapi.OptString{Set: false}
	var memBio = openapi.OptString{Set: false}

	dn := ""

	var usr openapi.User
	if m.MembershipStatus().IsLeft() {
		usr = a.ua.AdaptForLeft(m.User())
		dn = "Removed User"
	} else {
		usr = a.ua.Adapt(m.User())
		if p.HasDisplayName() {
			dn = p.DisplayName().ToString()
		}
		if p.HasBio() {
			memBio.Set = true
			memBio.Value = p.Bio().ToString()
		}
		if p.HasIDNumber() {
			memIDNum.Set = true
			memIDNum.Value = p.IDNumber().ToString()
		}
	}

	return openapi.Member{
		ID:               m.ID().Value(),
		User:             usr,
		Role:             a.adaptRole(m.Role()),
		MembershipStatus: memStatus,
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
