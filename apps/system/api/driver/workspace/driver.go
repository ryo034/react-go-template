package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	dbErr "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/error"
	"github.com/uptrace/bun"
)

type Driver interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (models.Workspaces, error)
	Create(ctx context.Context, exec bun.IDB, w *workspace.Workspace) (*models.Workspace, error)
	AddMember(ctx context.Context, exec bun.IDB, w *workspace.Workspace, m *member.Member) (*models.Member, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (p *driver) FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (models.Workspaces, error) {
	var ws models.Workspaces
	err := exec.
		NewSelect().
		Model(&ws).
		Relation("Detail").
		Join("JOIN members ms ON ms.workspace_id = ws.workspace_id").
		Where("ms.system_account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return ws, nil
}

func (p *driver) Create(ctx context.Context, exec bun.IDB, w *workspace.Workspace) (*models.Workspace, error) {
	d := w.Detail()
	m := &models.Workspace{
		WorkspaceID: w.ID().Value(),
	}
	md := &models.WorkspaceDetail{
		WorkspaceID: w.ID().Value(),
		Subdomain:   d.Subdomain().ToString(),
		Name:        d.Name().ToString(),
	}
	if _, err := exec.NewInsert().Model(m).Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := exec.NewInsert().Model(md).Exec(ctx); err != nil {
		if dbErr.IsDuplicateError(err) {
			return nil, domainErr.NewConflicted("workspace_details", "subdomain")
		}
		return nil, err
	}
	m.Detail = md
	return m, nil
}

func (p *driver) AddMember(ctx context.Context, exec bun.IDB, w *workspace.Workspace, m *member.Member) (*models.Member, error) {
	mm := &models.Member{
		MemberID:        m.ID().Value(),
		WorkspaceID:     w.ID().Value(),
		SystemAccountID: m.User().AccountID().Value(),
	}
	if _, err := exec.NewInsert().Model(mm).Exec(ctx); err != nil {
		return nil, err
	}

	mp := &models.MemberProfile{
		MemberID:       m.ID().Value(),
		MemberIDNumber: "",
		DisplayName:    m.DisplayName().ToString(),
	}
	if _, err := exec.NewInsert().Model(mp).Exec(ctx); err != nil {
		return nil, err
	}
	return mm, nil
}
