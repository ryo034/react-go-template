package me

import (
	"context"

	"github.com/google/uuid"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/id"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	invitationDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Find(ctx context.Context, exec bun.IDB, mID member.ID) (*models.Member, error)
	LastLogin(ctx context.Context, exec bun.IDB, mID member.ID) error
	FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*models.MemberLoginHistory, error)
	FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*models.SystemAccount, error)
	FindProfile(ctx context.Context, exec bun.IDB, aID account.ID) (*models.SystemAccount, error)
	FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error)
	UpdateName(ctx context.Context, exec bun.IDB, usr *user.User) error
	UpdateMemberProfile(ctx context.Context, exec bun.IDB, m *member.Member) (*member.Member, error)
}

type driver struct {
	invd invitationDr.Driver
}

func NewDriver(invd invitationDr.Driver) Driver {
	return &driver{invd}
}

func (d *driver) Find(ctx context.Context, exec bun.IDB, mID member.ID) (*models.Member, error) {
	mem := &models.Member{}
	err := exec.
		NewSelect().
		Model(mem).
		Relation("Profile").
		Relation("Role").
		Relation("Workspace").
		Relation("Workspace.Detail").
		Relation("SystemAccount").
		Relation("SystemAccount.AuthProviders").
		Relation("SystemAccount.Name").
		Relation("SystemAccount.Name.SystemAccountName").
		Relation("SystemAccount.Email").
		Relation("SystemAccount.Email.SystemAccountEmail").
		Relation("SystemAccount.PhoneNumber").
		Relation("SystemAccount.PhoneNumber.SystemAccountPhoneNumber").
		Relation("SystemAccount.PhotoEvent").
		Relation("SystemAccount.PhotoEvent.SystemAccountPhotoEvent").
		Relation("SystemAccount.PhotoEvent.SystemAccountPhotoEvent.Photo").
		Where("ms.member_id = ?", mID.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (d *driver) LastLogin(ctx context.Context, exec bun.IDB, mID member.ID) error {
	pkID, err := id.GenerateUUID()
	if err != nil {
		return err
	}
	m := &models.MemberLoginHistory{
		MemberLoginHistoryID: pkID.Value(),
		MemberID:             mID.Value(),
	}
	_, err = exec.NewInsert().Model(m).Exec(ctx)
	return err
}

func (d *driver) FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*models.MemberLoginHistory, error) {
	m := &models.MemberLoginHistory{}
	err := exec.
		NewSelect().
		Model(m).
		Relation("Member").
		Where("member.system_account_id = ?", aID.ToString()).
		Order("login_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d *driver) FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*models.SystemAccount, error) {
	sysAcc := &models.SystemAccount{}
	err := exec.
		NewSelect().
		Model(sysAcc).
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
		Where("sa.system_account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc, nil
}

func (d *driver) FindProfile(ctx context.Context, exec bun.IDB, aID account.ID) (*models.SystemAccount, error) {
	sysAcc := &models.SystemAccount{}
	err := exec.
		NewSelect().
		Model(sysAcc).
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
		Where("sa.system_account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc, nil
}

func (d *driver) FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error) {
	sysAcc := &models.SystemAccountEmail{}
	err := exec.
		NewSelect().
		Model(sysAcc).
		Relation("SystemAccount").
		Relation("SystemAccount.AuthProviders").
		Relation("SystemAccount.Name").
		Relation("SystemAccount.Name.SystemAccountName").
		Relation("SystemAccount.Email").
		Relation("SystemAccount.Email.SystemAccountEmail").
		Relation("SystemAccount.PhoneNumber").
		Relation("SystemAccount.PhoneNumber.SystemAccountPhoneNumber").
		Relation("SystemAccount.PhotoEvent").
		Relation("SystemAccount.PhotoEvent.SystemAccountPhotoEvent").
		Relation("SystemAccount.PhotoEvent.SystemAccountPhotoEvent.Photo").
		Where("saes.email = ?", email.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc.SystemAccount, nil
}

func (d *driver) UpdateName(ctx context.Context, exec bun.IDB, usr *user.User) error {
	sanID, err := uuid.NewV7()
	if err != nil {
		return err
	}
	mem := &models.SystemAccountName{
		SystemAccountNameID: sanID,
		SystemAccountID:     usr.AccountID().Value(),
		Name:                usr.Name().ToString(),
	}
	_, err = exec.NewInsert().Model(mem).Exec(ctx)
	return err
}

func (d *driver) UpdateMemberProfile(ctx context.Context, exec bun.IDB, m *member.Member) (*member.Member, error) {
	p := m.Profile()
	idNum := ""
	if p.IDNumber() != nil {
		idNum = p.IDNumber().ToString()
	}
	mem := &models.MemberProfile{
		MemberID:       m.ID().Value(),
		MemberIDNumber: idNum,
		DisplayName:    p.DisplayName().ToString(),
		Bio:            p.Bio().ToString(),
	}
	_, err := exec.NewUpdate().Model(mem).WherePK().Exec(ctx)
	return m, err
}
