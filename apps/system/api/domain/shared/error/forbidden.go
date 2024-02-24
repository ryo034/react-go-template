package error

type Forbidden struct {
	msg string
}

func NewForbidden(msg string) *Forbidden {
	return &Forbidden{msg}
}

func (u *Forbidden) Error() string {
	return u.msg
}

func (u *Forbidden) MessageKey() MessageKey {
	return ForbiddenMessageKey
}
