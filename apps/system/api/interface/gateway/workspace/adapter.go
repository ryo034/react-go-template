package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	Adapt(w *models.Workspace) (*workspace.Workspace, error)
	AdaptAll(ws []*models.Workspace) (workspace.Workspaces, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(w *models.Workspace) (*workspace.Workspace, error) {
	id := workspace.NewIDFromUUID(w.WorkspaceID)
	name, err := workspace.NewName(w.Detail.Name)
	if err != nil {
		return nil, err
	}
	d := workspace.NewDetail(name)
	return workspace.NewWorkspace(id, d), nil
}

func (a *adapter) AdaptAll(ws []*models.Workspace) (workspace.Workspaces, error) {
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
