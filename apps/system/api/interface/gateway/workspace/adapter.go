package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
)

type Adapter interface {
	Adapt(w *models.Workspace) (*workspace.Workspace, error)
	AdaptAll(ws []*models.Workspace) (workspace.Workspaces, error)
	AdaptInviter(w *models.Workspace, m *models.Member) (workspace.Inviter, error)
}

type adapter struct {
	ma member.Adapter
}

func NewAdapter(ma member.Adapter) Adapter {
	return &adapter{ma}
}

func (a *adapter) Adapt(w *models.Workspace) (*workspace.Workspace, error) {
	id := workspace.NewIDFromUUID(w.WorkspaceID)
	name, err := workspace.NewName(w.Detail.WorkspaceDetail.Name)
	if err != nil {
		return nil, err
	}
	sd, err := workspace.NewSubdomain(w.Detail.WorkspaceDetail.Subdomain)
	if err != nil {
		return nil, err
	}
	d := workspace.NewDetail(name, sd)
	return workspace.NewWorkspace(id, d), nil
}

func (a *adapter) AdaptAll(ws []*models.Workspace) (workspace.Workspaces, error) {
	if ws == nil {
		return nil, nil
	}
	mws := make([]*workspace.Workspace, 0, len(ws))
	for _, w := range ws {
		aw, err := a.Adapt(w)
		if err != nil {
			return nil, err
		}
		mws = append(mws, aw)
	}
	return workspace.NewWorkspaces(mws), nil
}

func (a *adapter) AdaptInviter(w *models.Workspace, m *models.Member) (workspace.Inviter, error) {
	mem, err := a.ma.Adapt(m)
	if err != nil {
		return workspace.Inviter{}, err
	}
	ws, err := a.Adapt(w)
	if err != nil {
		return workspace.Inviter{}, err
	}
	return workspace.NewInviter(mem, ws), nil
}
