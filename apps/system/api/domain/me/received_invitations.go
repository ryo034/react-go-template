package me

type ReceivedInvitations interface {
	Size() int
	AsSlice() []ReceivedInvitation
	IsEmpty() bool
	IsNotEmpty() bool
}

type workspaces struct {
	wrapped []ReceivedInvitation
}

func NewReceivedInvitations(wrapped []ReceivedInvitation) ReceivedInvitations {
	return &workspaces{wrapped}
}

func (ws *workspaces) IsEmpty() bool {
	return len(ws.wrapped) == 0
}

func (ws *workspaces) IsNotEmpty() bool {
	return !ws.IsEmpty()
}

func (ws *workspaces) Size() int {
	return len(ws.wrapped)
}

func (ws *workspaces) AsSlice() []ReceivedInvitation {
	return append(make([]ReceivedInvitation, 0, ws.Size()), ws.wrapped...)
}
