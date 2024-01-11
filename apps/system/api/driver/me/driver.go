package me

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Find(ctx context.Context, aID account.ID) (*models.Member, error)
}

type driver struct {
	db *bun.DB
}

func (d *driver) Find(ctx context.Context, aID account.ID) (*models.Member, error) {
	mem := &models.Member{}
	err := d.db.NewSelect().
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

func NewDriver(db *bun.DB) Driver {
	return &driver{db}
}
