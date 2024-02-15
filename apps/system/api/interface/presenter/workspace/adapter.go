package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(w *workspace.Workspace) openapi.Workspace
	AdaptAll(ws workspace.Workspaces) []openapi.Workspace
	AdaptInviter(i workspace.Inviter) openapi.Inviter
}

type adapter struct {
	ma member.Adapter
}

func NewAdapter(ma member.Adapter) Adapter {
	return &adapter{ma}
}

func (a *adapter) Adapt(w *workspace.Workspace) openapi.Workspace {
	if w == nil {
		return openapi.Workspace{}
	}
	d := w.Detail()
	return openapi.Workspace{
		WorkspaceId: w.ID().Value(),
		Name:        d.Name().ToString(),
		Subdomain:   d.Subdomain().ToString(),
	}
}

func (a *adapter) AdaptAll(ws workspace.Workspaces) []openapi.Workspace {
	if ws == nil {
		return nil
	}
	res := make([]openapi.Workspace, 0, ws.Size())
	for _, w := range ws.AsSlice() {
		res = append(res, a.Adapt(w))
	}
	return res
}

func (a *adapter) AdaptInviter(i workspace.Inviter) openapi.Inviter {
	return openapi.Inviter{
		Member:    a.ma.Adapt(i.Member),
		Workspace: a.Adapt(i.Workspace()),
	}
}
