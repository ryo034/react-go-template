package workspace

type Workspace struct {
	id     ID
	detail Detail
}

func NewWorkspace(id ID, detail Detail) *Workspace {
	return &Workspace{id, detail}
}

func (w *Workspace) ID() ID {
	return w.id
}

func (w *Workspace) Detail() Detail {
	return w.detail
}
