package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	FindAll(ctx context.Context, exec bun.IDB, aID account.ID) (models.Workspaces, error)
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
