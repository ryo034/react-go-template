package member

type InvitedMembers interface {
	Size() int
	AsSlice() []*InvitedMember
	IsEmpty() bool
	IsNotEmpty() bool
}

type invitedMembers struct {
	wrapped []*InvitedMember
}

func NewInvitedMembers(wrapped []*InvitedMember) InvitedMembers {
	return &invitedMembers{wrapped}
}

func (ms *invitedMembers) IsEmpty() bool {
	return len(ms.wrapped) == 0
}

func (ms *invitedMembers) IsNotEmpty() bool {
	return !ms.IsEmpty()
}

func (ms *invitedMembers) Size() int {
	return len(ms.wrapped)
}

func (ms *invitedMembers) AsSlice() []*InvitedMember {
	return append(make([]*InvitedMember, 0, ms.Size()), ms.wrapped...)
}
