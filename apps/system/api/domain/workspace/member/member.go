package member

import "github.com/ryo034/react-go-template/apps/system/api/domain/user"

type Member struct {
	id          ID
	u           *user.User
	displayName DisplayName
	idNumber    *IDNumber
}

func NewMember(id ID, u *user.User, displayName DisplayName, idNumber *IDNumber) *Member {
	return &Member{id, u, displayName, idNumber}
}

func NewMemberFromUser(u *user.User, displayName DisplayName) (*Member, error) {
	var dn DisplayName
	var err error
	if displayName.IsNotEmpty() {
		dn = displayName
	} else {
		dn, err = NewDisplayName(u.Name().ToString())
		if err != nil {
			return nil, err
		}
	}
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}
	return &Member{id, u, dn, nil}, nil
}

func (w *Member) ID() ID {
	return w.id
}

func (w *Member) User() *user.User {
	return w.u
}

func (w *Member) DisplayName() DisplayName {
	return w.displayName
}

func (w *Member) IDNumber() *IDNumber {
	return w.idNumber
}

func (w *Member) HasIDNumber() bool {
	return w.idNumber != nil
}
