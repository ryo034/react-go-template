package member

import (
	"fmt"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
)

type Member struct {
	id               ID
	u                *user.User
	profile          Profile
	role             Role
	membershipStatus MembershipStatus
}

func NewMember(id ID, u *user.User, profile Profile, role Role, participationStatus MembershipStatus) *Member {
	dn := profile.DisplayName()
	if profile.DisplayName() == nil {
		dn = NewDisplayName(u.Name().ToString())
	}
	profile.displayName = dn
	return &Member{id, u, profile, role, participationStatus}
}

func GenerateMember(u *user.User) (*Member, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewMember(id, u, NewEmptyProfile(), RoleMember, MembershipStatusActive), nil
}

func GenerateAsWorkspaceOwner(u *user.User, dn *DisplayName) (*Member, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewMember(id, u, NewProfile(dn, nil, NewAsEmptyBio()), RoleOwner, MembershipStatusActive), nil
}

func (m *Member) ID() ID {
	return m.id
}

func (m *Member) User() *user.User {
	return m.u
}

func (m *Member) Profile() Profile {
	return m.profile
}

func (m *Member) Role() Role {
	return m.role
}

func (m *Member) MembershipStatus() MembershipStatus {
	return m.membershipStatus
}

// UpdateProfile updates the profile of the member
// if the displayName is nil, it will be set to the user's name
func (m *Member) UpdateProfile(profile Profile) (*Member, error) {
	if m == nil {
		return nil, domainErr.NewUnauthenticated("Not joined")
	}
	if m.membershipStatus.IsLeft() {
		return nil, domainErr.NewGone(fmt.Sprintf("MemberID %s", m.ID().ToString()))
	}
	if profile.DisplayName() == nil {
		profile.displayName = NewDisplayName(m.u.Name().ToString())
	}
	m.profile = profile
	return m, nil
}

func (m *Member) UpdateUser(u *user.User) *Member {
	m.u = u
	return m
}

func (m *Member) UpdateRole(role Role) (*Member, error) {
	if m.membershipStatus.IsLeft() {
		return nil, domainErr.NewGone(fmt.Sprintf("MemberID %s", m.ID().ToString()))
	}
	if role == RoleOwner {
		return nil, domainErr.NewForbidden("cannot change the role to owner")
	}
	if role == m.role {
		return nil, domainErr.NewBadRequest("the role is already the same")
	}
	if m.role.IsOwner() {
		return nil, domainErr.NewForbidden("cannot change the role")
	}
	m.role = role
	return m, nil
}

func (m *Member) ValidateCanLeave() error {
	if m.MembershipStatus().IsLeft() {
		return domainErr.NewForbidden("already left")
	}
	return nil
}
