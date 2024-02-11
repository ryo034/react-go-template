package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	memberPresenter "github.com/ryo034/react-go-template/apps/system/api/interface/presenter/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	workspaceUc "github.com/ryo034/react-go-template/apps/system/api/usecase/workspace"
)

func NewPresenter(wa Adapter, ma memberPresenter.Adapter) workspaceUc.OutputPort {
	return &presenter{wa, ma}
}

type presenter struct {
	wa Adapter
	ma memberPresenter.Adapter
}

func (p *presenter) Create(w *workspace.Workspace) *openapi.Workspace {
	res := p.wa.Adapt(w)
	return &res
}

func (p *presenter) FindAllMembers(ms member.Members) *openapi.Members {
	res := p.ma.AdaptAll(ms)
	return &res
}
