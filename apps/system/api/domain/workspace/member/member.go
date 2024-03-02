package member

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
)

type Member struct {
	id      ID
	u       *user.User
	profile Profile
	role    Role
}

func NewMember(id ID, u *user.User, profile Profile, role Role) *Member {
	dn := profile.DisplayName()
	if profile.DisplayName() == nil {
		dn = NewDisplayName(u.Name().ToString())
	}
	profile.displayName = dn
	return &Member{id, u, profile, role}
}

func GenerateMember(u *user.User) (*Member, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewMember(id, u, NewEmptyProfile(), RoleMember), nil
}

func GenerateAsWorkspaceOwner(u *user.User, dn *DisplayName) (*Member, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return NewMember(id, u, NewProfile(dn, nil, NewAsEmptyBio()), RoleOwner), nil
}

func (w *Member) ID() ID {
	return w.id
}

func (w *Member) User() *user.User {
	return w.u
}

func (w *Member) Profile() Profile {
	return w.profile
}

func (w *Member) Role() Role {
	return w.role
}

// UpdateProfile updates the profile of the member
// if the displayName is nil, it will be set to the user's name
func (w *Member) UpdateProfile(profile Profile) *Member {
	if profile.DisplayName() == nil {
		profile.displayName = NewDisplayName(w.u.Name().ToString())
	}
	w.profile = profile
	return w
}

func (w *Member) UpdateUser(u *user.User) *Member {
	w.u = u
	return w
}

func (w *Member) UpdateRole(role Role) *Member {
	w.role = role
	return w
}
