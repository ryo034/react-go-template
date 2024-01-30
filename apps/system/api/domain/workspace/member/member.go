package member

import "github.com/ryo034/react-go-template/apps/system/api/domain/user"

type Member struct {
	id          ID
	u           *user.User
	displayName *DisplayName
	idNumber    *IDNumber
}

func NewMember(id ID, u *user.User, displayName *DisplayName, idNumber *IDNumber) *Member {
	return &Member{id, u, displayName, idNumber}
}

func (w *Member) ID() ID {
	return w.id
}

func (w *Member) User() *user.User {
	return w.u
}

func (w *Member) DisplayName() *DisplayName {
	return w.displayName
}

func (w *Member) HasDisplayName() bool {
	return w.displayName != nil
}

func (w *Member) IDNumber() *IDNumber {
	return w.idNumber
}

func (w *Member) HasIDNumber() bool {
	return w.idNumber != nil
}
