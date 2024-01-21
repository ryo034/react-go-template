package member

import "github.com/ryo034/react-go-template/apps/system/api/domain/user"

type Member struct {
	id       ID
	u        *user.User
	profile  *Profile
	idNumber string
}

func NewMember(id ID, u *user.User, profile *Profile, idNumber string) *Member {
	return &Member{id, u, profile, idNumber}
}

func (w *Member) ID() ID {
	return w.id
}

func (w *Member) User() *user.User {
	return w.u
}

func (w *Member) Profile() *Profile {
	return w.profile
}

func (w *Member) IDNumber() string {
	return w.idNumber
}
