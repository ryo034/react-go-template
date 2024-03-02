package member

type Role string

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleGuest  Role = "guest"
)

func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

func (r Role) IsMember() bool {
	return r == RoleMember
}

func (r Role) IsGuest() bool {
	return r == RoleGuest
}
