package invitation

type Invitations interface {
	Size() int
	AsSlice() []*Invitation
	IsEmpty() bool
	IsNotEmpty() bool
}

type workspaces struct {
	wrapped []*Invitation
}

func NewInvitations(wrapped []*Invitation) Invitations {
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

func (ws *workspaces) AsSlice() []*Invitation {
	return append(make([]*Invitation, 0, ws.Size()), ws.wrapped...)
}
