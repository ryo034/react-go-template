package workspace

type Workspaces interface {
	Size() int
	AsSlice() []*Workspace
	IsEmpty() bool
	IsNotEmpty() bool
}

type workspaces struct {
	wrapped []*Workspace
}

func NewWorkspaces(wrapped []*Workspace) Workspaces {
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

func (ws *workspaces) AsSlice() []*Workspace {
	return append(make([]*Workspace, 0, ws.Size()), ws.wrapped...)
}
