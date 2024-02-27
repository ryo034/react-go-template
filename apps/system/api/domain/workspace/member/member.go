package member

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
)

type Member struct {
	id      ID
	u       *user.User
	profile Profile
}

func NewMember(id ID, u *user.User, profile Profile) *Member {
	dn := profile.DisplayName()
	if profile.DisplayName() == nil {
		dn = NewDisplayName(u.Name().ToString())
	}
	profile.displayName = dn
	return &Member{id, u, profile}
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

// UpdateProfile updates the profile of the member
// if the displayName is nil, it will be set to the user's name
func (w *Member) UpdateProfile(profile Profile) *Member {
	if profile.DisplayName() == nil {
		profile.displayName = NewDisplayName(w.u.Name().ToString())
	}
	w.profile = profile
	return w
}
