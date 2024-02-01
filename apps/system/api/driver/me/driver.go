package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Find(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Member, error)
	FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*models.SystemAccount, error)
	Update(ctx context.Context, exec bun.IDB, m *me.Me) (*models.Member, error)
	UpdateName(ctx context.Context, exec bun.IDB, aID account.ID, name account.Name) (*models.SystemAccount, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) Find(ctx context.Context, exec bun.IDB, aID account.ID, wID workspace.ID) (*models.Member, error) {
	mem := &models.Member{}
	err := exec.
		NewSelect().
		Model(mem).
		Relation("Profile").
		Relation("SystemAccount").
		Relation("SystemAccount.Profile").
		Relation("SystemAccount.PhoneNumber").
		Where("m.system_account_id = ?", aID.ToString()).
		Where("m.workspace_id = ?", wID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (d *driver) FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*models.SystemAccount, error) {
	sysAcc := &models.SystemAccount{}
	err := exec.
		NewSelect().
		Model(sysAcc).
		Relation("Profile").
		Relation("PhoneNumber").
		Where("sa.system_account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc, nil
}

func (d *driver) Update(ctx context.Context, exec bun.IDB, m *me.Me) (*models.Member, error) {
	mem := &models.Member{}
	_, err := exec.
		NewUpdate().
		Model(mem).
		Set("updated_at = now()").
		Where("id = ?", m.Member().ID().ToString()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (d *driver) UpdateName(ctx context.Context, exec bun.IDB, aID account.ID, name account.Name) (*models.SystemAccount, error) {
	res, err := d.FindBeforeOnboard(ctx, exec, aID)
	if err != nil {
		return nil, err
	}
	res.Profile.Name = name.ToString()
	m := models.SystemAccountProfile{
		SystemAccountID: aID.Value(),
		Name:            name.ToString(),
		Email:           res.Profile.Email,
	}
	_, err = exec.
		NewUpdate().
		Model(&m).
		Where("system_account_id = ?", aID.ToString()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
