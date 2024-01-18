package membership

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Find(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Membership, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) Find(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Membership, error) {
	mem := &models.Membership{}
	err := exec.
		NewSelect().
		Model(mem).
		Relation("Member").
		Relation("Workspace").
		Relation("Workspace.Detail").
		Relation("Member.Profile").
		Relation("Member.SystemAccount").
		Relation("Member.SystemAccount.Profile").
		Relation("Member.SystemAccount.PhoneNumber").
		Where("member.system_account_id = ?", aID.ToString()).
		Where("mship.workspace_id = ?", wID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}
