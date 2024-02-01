package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

func NewPresenter(wa Adapter) workspaceUc.OutputPort {
	return &presenter{wa}
}

type presenter struct {
	wa Adapter
}

func (p *presenter) Create(w *workspace.Workspace) *openapi.Workspace {
	res := p.wa.Adapt(w)
	return &res
}
