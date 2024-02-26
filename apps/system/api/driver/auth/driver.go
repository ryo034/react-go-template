package auth

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Create(ctx context.Context, exec bun.IDB, aID account.ID, email account.Email, ap *provider.Provider) (*models.SystemAccount, error)
	Find(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (p *driver) Create(ctx context.Context, exec bun.IDB, aID account.ID, email account.Email, ap *provider.Provider) (*models.SystemAccount, error) {
	sa := models.SystemAccount{
		SystemAccountID: aID.Value(),
	}
	_, err := exec.
		NewInsert().
		Model(&sa).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	sap := models.SystemAccountProfile{
		SystemAccountID: aID.Value(),
		Name:            "",
	}
	_, err = exec.
		NewInsert().
		Model(&sap).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	sape := models.SystemAccountEmail{
		SystemAccountID: aID.Value(),
		Email:           email.ToString(),
	}
	if _, err = exec.
		NewInsert().
		Model(&sape).
		Exec(ctx); err != nil {
		return nil, err
	}

	prb := "firebase"
	switch ap.ProvidedBy() {
	case provider.ProvidedByFirebase:
		prb = "firebase"
	}

	prv := "email"
	switch ap.Kind() {
	case provider.Email:
		prv = "email"
	case provider.Google:
		prb = "google"
	}

	apm := models.AuthProvider{
		AuthProviderID:  ap.ID().Value(),
		SystemAccountID: aID.Value(),
		Provider:        prv,
		ProvidedBy:      prb,
	}
	if _, err = exec.
		NewInsert().
		Model(&apm).
		Exec(ctx); err != nil {
		return nil, err
	}

	sa.Profile = &sap
	sa.Emails = append(sa.Emails, &sape)
	sa.AuthProviders = append(sa.AuthProviders, &apm)

	return &sa, err
}

func (p *driver) Find(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error) {
	sap := &models.SystemAccountEmail{}
	err := exec.
		NewSelect().
		Model(sap).
		Where("email = ?", email.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	sa := &models.SystemAccount{}
	err = exec.
		NewSelect().
		Model(sa).
		Relation("Profile").
		Relation("AuthProviders").
		Relation("Emails").
		Relation("PhoneNumbers").
		Where("sa.system_account_id = ?", sap.SystemAccountID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sa, nil
}
