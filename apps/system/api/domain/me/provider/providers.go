package provider

type Providers interface {
	Size() int
	AsSlice() []*Provider
	IsEmpty() bool
	IsNotEmpty() bool
	FindByKind(k Kind) *Provider
}

type workspaces struct {
	wrapped []*Provider
}

func NewProviders(wrapped []*Provider) Providers {
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

func (ws *workspaces) AsSlice() []*Provider {
	return append(make([]*Provider, 0, ws.Size()), ws.wrapped...)
}

func (ws *workspaces) FindByKind(k Kind) *Provider {
	for _, w := range ws.wrapped {
		if w.Kind() == k {
			return w
		}
	}
	return nil
}
