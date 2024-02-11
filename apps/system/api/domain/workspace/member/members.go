package member

type Members interface {
	Size() int
	AsSlice() []*Member
	IsEmpty() bool
	IsNotEmpty() bool
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
