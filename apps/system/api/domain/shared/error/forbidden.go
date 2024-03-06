package error

type Forbidden struct {
	msg string
}

func NewForbidden(msg string) *Forbidden {
	return &Forbidden{msg}
}

func (e *Forbidden) Error() string {
	return e.msg
}

func (e *Forbidden) MessageKey() MessageKey {
	return ForbiddenMessageKey
}
