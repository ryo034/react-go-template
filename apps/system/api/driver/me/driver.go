package me

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"

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
	FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*models.Account, error)
	FindProfile(ctx context.Context, exec bun.IDB, aID account.ID) (*models.Account, error)
	FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Account, error)
	UpdateName(ctx context.Context, exec bun.IDB, usr *user.User) error
	UpdateMemberProfile(ctx context.Context, exec bun.IDB, m *member.Member) (*member.Member, error)
	UpdateProfilePhoto(ctx context.Context, exec bun.IDB, aID account.ID, photo *media.UploadPhoto) error
	RemoveProfilePhoto(ctx context.Context, exec bun.IDB, aID account.ID) error
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
		Relation("Role.MemberRole").
		Relation("Workspace").
		Relation("Workspace.Detail").
		Relation("Account").
		Relation("Account.AuthProviders").
		Relation("Account.Name").
		Relation("Account.Name.AccountName").
		Relation("Account.Email").
		Relation("Account.Email.AccountEmail").
		Relation("Account.PhoneNumber").
		Relation("Account.PhoneNumber.AccountPhoneNumber").
		Relation("Account.PhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent.Photo").
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
	_, err = exec.NewDelete().Model(&models.MemberLatestLoginHistory{}).Where("member_id = ?", mID.Value()).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.MemberLoginHistory{
		MemberLoginHistoryID: pkID.Value(),
		MemberID:             mID.Value(),
	}).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.MemberLatestLoginHistory{
		MemberLoginHistoryID: pkID.Value(),
		MemberID:             mID.Value(),
	}).Exec(ctx)

	return err
}

func (d *driver) FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*models.MemberLoginHistory, error) {
	m := &models.MemberLoginHistory{}
	err := exec.
		NewSelect().
		Model(m).
		Relation("Member").
		Where("member.account_id = ?", aID.ToString()).
		Order("login_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d *driver) FindBeforeOnboard(ctx context.Context, exec bun.IDB, aID account.ID) (*models.Account, error) {
	sysAcc := &models.Account{}
	err := exec.
		NewSelect().
		Model(sysAcc).
		Relation("AuthProviders").
		Relation("Name").
		Relation("Name.AccountName").
		Relation("Email").
		Relation("Email.AccountEmail").
		Relation("PhoneNumber").
		Relation("PhoneNumber.AccountPhoneNumber").
		Relation("PhotoEvent").
		Relation("PhotoEvent.AccountPhotoEvent").
		Relation("PhotoEvent.AccountPhotoEvent.Photo").
		Where("sa.account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc, nil
}

func (d *driver) FindProfile(ctx context.Context, exec bun.IDB, aID account.ID) (*models.Account, error) {
	sysAcc := &models.Account{}
	err := exec.
		NewSelect().
		Model(sysAcc).
		Relation("AuthProviders").
		Relation("Name").
		Relation("Name.AccountName").
		Relation("Email").
		Relation("Email.AccountEmail").
		Relation("PhoneNumber").
		Relation("PhoneNumber.AccountPhoneNumber").
		Relation("PhotoEvent").
		Relation("PhotoEvent.AccountPhotoEvent").
		Relation("PhotoEvent.AccountPhotoEvent.Photo").
		Where("sa.account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc, nil
}

func (d *driver) FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.Account, error) {
	sysAcc := &models.AccountEmail{}
	err := exec.
		NewSelect().
		Model(sysAcc).
		Relation("Account").
		Relation("Account.AuthProviders").
		Relation("Account.Name").
		Relation("Account.Name.AccountName").
		Relation("Account.Email").
		Relation("Account.Email.AccountEmail").
		Relation("Account.PhoneNumber").
		Relation("Account.PhoneNumber.AccountPhoneNumber").
		Relation("Account.PhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent").
		Relation("Account.PhotoEvent.AccountPhotoEvent.Photo").
		Where("ae.email = ?", email.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc.Account, nil
}

func (d *driver) UpdateName(ctx context.Context, exec bun.IDB, usr *user.User) error {
	sanID, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = exec.NewDelete().Model(&models.AccountLatestName{}).Where("account_id = ?", usr.AccountID().Value()).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.AccountName{
		AccountNameID: sanID,
		AccountID:     usr.AccountID().Value(),
		Name:          usr.Name().ToString(),
	}).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.AccountLatestName{
		AccountNameID: sanID,
		AccountID:     usr.AccountID().Value(),
	}).Exec(ctx)
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

func (d *driver) UpdateProfilePhoto(ctx context.Context, exec bun.IDB, aID account.ID, photo *media.UploadPhoto) error {
	peID, err := id.GenerateUUID()
	if err != nil {
		return err
	}
	_, err = exec.NewDelete().Model(&models.AccountLatestPhotoEvent{}).Where("account_id = ?", aID.ToString()).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.AccountPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
		EventType:           "upload",
	}).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.AccountLatestPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
	}).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.AccountPhoto{
		AccountPhotoEventID: peID.Value(),
		PhotoID:             photo.ID().Value(),
		HostingTo:           photo.HostingTo().String(),
	}).Exec(ctx)
	return err
}

func (d *driver) RemoveProfilePhoto(ctx context.Context, exec bun.IDB, aID account.ID) error {
	peID, err := id.GenerateUUID()
	if err != nil {
		return err
	}
	_, err = exec.NewDelete().Model(&models.AccountLatestPhotoEvent{}).Where("account_id = ?", aID.ToString()).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.AccountPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
		EventType:           "remove",
	}).Exec(ctx)
	_, err = exec.NewInsert().Model(&models.AccountLatestPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
	}).Exec(ctx)
	return err
}
