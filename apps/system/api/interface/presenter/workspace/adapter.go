package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(w *workspace.Workspace) openapi.Workspace
	AdaptAll(ws workspace.Workspaces) []openapi.Workspace
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(w *workspace.Workspace) openapi.Workspace {
	if w == nil {
		return openapi.Workspace{}
	}
	d := w.Detail()
	return openapi.Workspace{
		WorkspaceId: w.ID().ToFriendlyString(),
		Name:        d.Name().ToString(),
		Subdomain:   d.Subdomain().ToString(),
	}
}

func (a *adapter) AdaptAll(ws workspace.Workspaces) []openapi.Workspace {
	res := make([]openapi.Workspace, 0, ws.Size())
	for _, w := range ws.AsSlice() {
		res = append(res, a.Adapt(w))
	}
	return res
}
