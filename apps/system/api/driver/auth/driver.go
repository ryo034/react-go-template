package auth

import (
	"context"

	"github.com/google/uuid"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Create(ctx context.Context, exec bun.IDB, usr *user.User, ap *provider.Provider) (*models.SystemAccount, error)
	Find(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error)
	FindAccountIDByAuthProviderUID(ctx context.Context, exec bun.IDB, apUID provider.UID) (uuid.UUID, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (p *driver) Create(ctx context.Context, exec bun.IDB, usr *user.User, ap *provider.Provider) (*models.SystemAccount, error) {
	sa := models.SystemAccount{
		SystemAccountID: usr.AccountID().Value(),
	}
	_, err := exec.
		NewInsert().
		Model(&sa).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	na := ""
	if usr.HasName() {
		na = usr.Name().ToString()
	}

	sap := models.SystemAccountProfile{
		SystemAccountID: usr.AccountID().Value(),
		Name:            na,
	}
	_, err = exec.
		NewInsert().
		Model(&sap).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	sape := models.SystemAccountEmail{
		SystemAccountID: usr.AccountID().Value(),
		Email:           usr.Email().ToString(),
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
		prv = "google"
	}

	apm := models.AuthProvider{
		AuthProviderID:  ap.ID().Value(),
		SystemAccountID: usr.AccountID().Value(),
		Provider:        prv,
		ProvidedBy:      prb,
		ProviderUID:     ap.UID().ToString(),
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

func (p *driver) FindAccountIDByAuthProviderUID(ctx context.Context, exec bun.IDB, apUID provider.UID) (uuid.UUID, error) {
	apm := &models.AuthProvider{}
	err := exec.
		NewSelect().
		Model(apm).
		Where("provider_uid = ?", apUID.ToString()).
		Scan(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}
	return apm.SystemAccountID, nil
}
