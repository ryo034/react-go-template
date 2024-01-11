package member

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
)

type Member struct {
	u                 *user.User
	idNumber          string
	membershipPeriods MembershipPeriods
}

func NewMember(u *user.User, idNumber string, membershipPeriods MembershipPeriods) *Member {
	return &Member{u, idNumber, membershipPeriods}
}

func (m *Member) User() *user.User {
	return m.u
}

func (m *Member) IDNumber() string {
	return m.idNumber
}

func (m *Member) MembershipPeriods() MembershipPeriods {
	return m.membershipPeriods
}
