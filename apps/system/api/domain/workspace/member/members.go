package member

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

type Members interface {
	Size() int
	AsSlice() []*Member
	IsEmpty() bool
	IsNotEmpty() bool
	Exist(email account.Email) bool
}

type members struct {
	wrapped []*Member
}

func NewMembers(wrapped []*Member) Members {
	return &members{wrapped}
}

func (ms *members) IsEmpty() bool {
	return len(ms.wrapped) == 0
}

func (ms *members) IsNotEmpty() bool {
	return !ms.IsEmpty()
}

func (ms *members) Size() int {
	return len(ms.wrapped)
}

func (ms *members) AsSlice() []*Member {
	return append(make([]*Member, 0, ms.Size()), ms.wrapped...)
}

func (ms *members) Exist(email account.Email) bool {
	for _, m := range ms.wrapped {
		if m.User().Email().ToString() == email.ToString() {
			return true
		}
	}
	return false
}
