package auth

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Create(ctx context.Context, exec bun.IDB, aID account.ID, email account.Email) (*models.SystemAccount, error)
	Find(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (p *driver) Create(ctx context.Context, exec bun.IDB, aID account.ID, email account.Email) (*models.SystemAccount, error) {
	sa := &models.SystemAccount{
		SystemAccountID: aID.Value(),
		Profile: &models.SystemAccountProfile{
			SystemAccountID: aID.Value(),
			Name:            "",
			Email:           email.ToString(),
		},
	}
	_, err := exec.
		NewInsert().
		Model(sa).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = exec.
		NewInsert().
		Model(sa.Profile).
		Exec(ctx)

	return sa, err
}

func (p *driver) Find(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error) {
	sa := &models.SystemAccount{}
	err := exec.
		NewSelect().
		Model(sa).
		Relation("Profile").
		Relation("PhoneNumber").
		Where("profile.email = ?", email.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sa, nil
}
