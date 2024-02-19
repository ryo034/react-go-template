package error

type Unauthenticated struct {
}

func NewUnauthenticated() *Unauthenticated {
	return &Unauthenticated{}
}

func (u *Unauthenticated) Error() string {
	return "Unauthenticated"
}

func (u *Unauthenticated) MessageKey() MessageKey {
	return UnauthenticatedMessageKey
}
