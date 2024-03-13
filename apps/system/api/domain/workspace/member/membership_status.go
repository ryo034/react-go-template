package member

import (
	"fmt"

	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type MembershipStatus string

const InvalidMemberMembershipStatus domainError.MessageKey = "invalid.member.membership_status"

const (
	MembershipStatusActive MembershipStatus = "active"
	MembershipStatusLeave  MembershipStatus = "leave"
)

func NewMembershipStatus(s string) (MembershipStatus, error) {
	errs := validation.NewErrors()
	switch s {
	case MembershipStatusActive.ToString():
		return MembershipStatusActive, nil
	case MembershipStatusLeave.ToString():
		return MembershipStatusLeave, nil
	default:
		errs.Append(InvalidMemberMembershipStatus, fmt.Sprintf("invalid membership status: %s", s))
		return "", errs
	}
}

func (ms MembershipStatus) ToString() string {
	return string(ms)
}

func (ms MembershipStatus) IsActive() bool {
	return ms == MembershipStatusActive
}

func (ms MembershipStatus) IsLeft() bool {
	return ms == MembershipStatusLeave
}
