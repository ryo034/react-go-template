package error

type NotBelong struct {
}

func NewNotBelong() *NotBelong {
	return &NotBelong{}
}

func (u *NotBelong) Error() string {
	return "NotBelong"
}

func (u *NotBelong) MessageKey() MessageKey {
	return NotBelongMessageKey
}

func (u *NotBelong) Code() string {
	return "400-" + string(NotBelongCodeKey)
}
