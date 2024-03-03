package member

import (
	"fmt"

	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type Role string

const InvalidMemberRole domainError.MessageKey = "invalid.member.role"

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleGuest  Role = "guest"
)

func (r Role) ToString() string {
	return string(r)
}

func NewRole(s string) (Role, error) {
	errs := validation.NewErrors()
	switch s {
	case RoleOwner.ToString():
		return RoleOwner, nil
	case RoleAdmin.ToString():
		return RoleAdmin, nil
	case RoleMember.ToString():
		return RoleMember, nil
	case RoleGuest.ToString():
		return RoleGuest, nil
	default:
		errs.Append(InvalidMemberRole, fmt.Sprintf("invalid role: %s", s))
		return "", errs
	}
}

func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

func (r Role) IsMember() bool {
	return r == RoleMember
}

func (r Role) IsGuest() bool {
	return r == RoleGuest
}
