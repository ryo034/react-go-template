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
	_, err := exec.NewInsert().Model(&sa).Exec(ctx)
	if err != nil {
		return nil, err
	}

	if usr.HasName() {
		if _, err = exec.NewDelete().
			Model(&models.SystemAccountLatestName{}).
			Where("system_account_id = ?", usr.AccountID().Value()).
			Exec(ctx); err != nil {
			return nil, err
		}

		na := usr.Name().ToString()
		sanID, err := uuid.NewV7()
		if err != nil {
			return nil, err
		}
		sap := &models.SystemAccountName{
			SystemAccountNameID: sanID,
			SystemAccountID:     usr.AccountID().Value(),
			Name:                na,
		}
		if _, err = exec.NewInsert().Model(sap).Exec(ctx); err != nil {
			return nil, err
		}

		saln := &models.SystemAccountLatestName{
			SystemAccountNameID: sanID,
			SystemAccountID:     usr.AccountID().Value(),
		}
		if _, err = exec.NewInsert().Model(saln).Exec(ctx); err != nil {
			return nil, err
		}

		sa.Name = saln
	}

	// Update Email
	if _, err = exec.NewDelete().
		Model(&models.SystemAccountLatestEmail{}).
		Where("system_account_id = ?", usr.AccountID().Value()).
		Exec(ctx); err != nil {
		return nil, err
	}

	saemID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	saem := &models.SystemAccountEmail{
		SystemAccountEmailID: saemID,
		SystemAccountID:      usr.AccountID().Value(),
		Email:                usr.Email().ToString(),
	}
	if _, err = exec.NewInsert().Model(saem).Exec(ctx); err != nil {
		return nil, err
	}

	salem := &models.SystemAccountLatestEmail{
		SystemAccountLatestEmailID: saemID,
		SystemAccountID:            usr.AccountID().Value(),
	}
	if _, err = exec.NewInsert().Model(salem).Exec(ctx); err != nil {
		return nil, err
	}

	sa.Email = salem

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
		Relation("AuthProviders").
		Relation("Name").
		Relation("Name.SystemAccountName").
		Relation("Email").
		Relation("Email.SystemAccountEmail").
		Relation("PhoneNumber").
		Relation("PhoneNumber.SystemAccountPhoneNumber").
		Relation("PhotoEvent").
		Relation("PhotoEvent.SystemAccountPhotoEvent").
		Relation("PhotoEvent.SystemAccountPhotoEvent.Photo").
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
