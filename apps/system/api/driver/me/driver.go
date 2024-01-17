package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Find(ctx context.Context, exec bun.IDB, aID account.ID) (*models.Member, error)
	Update(ctx context.Context, exec bun.IDB, m *me.Me) (*models.Member, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) Find(ctx context.Context, exec bun.IDB, aID account.ID) (*models.Member, error) {
	mem := &models.Member{}
	err := exec.
		NewSelect().
		Model(mem).
		Relation("Profile").
		Relation("SystemAccount").
		Relation("SystemAccount.Profile").
		Relation("SystemAccount.PhoneNumber").
		Where("m.system_account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (d *driver) Update(ctx context.Context, exec bun.IDB, m *me.Me) (*models.Member, error) {
	mem := &models.Member{}
	_, err := exec.
		NewUpdate().
		Model(mem).
		Set("updated_at = now()").
		Where("id = ?", m.Member().User().AccountID().ToString()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}
