package error

type Unauthenticated struct {
	msg string
}

func NewUnauthenticated(msg string) *Unauthenticated {
	return &Unauthenticated{msg}
}

func (u *Unauthenticated) Error() string {
	return u.msg
}

func (u *Unauthenticated) MessageKey() MessageKey {
	return UnauthenticatedMessageKey
}
