package me

import (
	"context"
	"database/sql"

	workspaceDr "github.com/ryo034/react-go-template/apps/system/api/driver/workspace"

	"github.com/go-faster/errors"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

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
	wd   workspaceDr.Driver
}

func NewDriver(invd invitationDr.Driver, wd workspaceDr.Driver) Driver {
	return &driver{invd, wd}
}

func (d *driver) Find(ctx context.Context, exec bun.IDB, mID member.ID) (*models.Member, error) {
	mem := &models.Member{}
	err := exec.
		NewSelect().
		Model(mem).
		Relation("Profile").
		Relation("Profile.MemberProfile").
		Relation("Role").
		Relation("Role.MemberRole").
		Relation("MembershipEvent").
		Relation("MembershipEvent.MembershipEvent").
		Relation("Workspace").
		Relation("Workspace.Detail").
		Relation("Workspace.Detail.WorkspaceDetail").
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
	if _, err = exec.NewDelete().Model(&models.MemberLatestLoginHistory{}).Where("member_id = ?", mID.Value()).
		Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.MemberLoginHistory{
		MemberLoginHistoryID: pkID.Value(),
		MemberID:             mID.Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.MemberLatestLoginHistory{
		MemberLoginHistoryID: pkID.Value(),
		MemberID:             mID.Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (d *driver) FindLastLogin(ctx context.Context, exec bun.IDB, aID account.ID) (*models.MemberLoginHistory, error) {
	m := &models.MemberLoginHistory{}
	err := exec.
		NewSelect().
		Model(m).
		Relation("Member").
		Relation("Member.MembershipEvent").
		Relation("Member.MembershipEvent.MembershipEvent").
		Where("ms.account_id = ?", aID.ToString()).
		Where("ms__lmshi__mshi.event_type = ?", "join").
		Order("mllhs.login_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErr.NewNoSuchData("No login history")
		}
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
	if usr.Name() == nil {
		return nil
	}
	sanID, err := uuid.NewV7()
	if err != nil {
		return err
	}
	if _, err = exec.NewDelete().Model(&models.AccountLatestName{}).Where("account_id = ?", usr.AccountID().Value()).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.AccountName{
		AccountNameID: sanID,
		AccountID:     usr.AccountID().Value(),
		Name:          usr.Name().ToString(),
	}).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.AccountLatestName{
		AccountNameID: sanID,
		AccountID:     usr.AccountID().Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	return err
}

func (d *driver) UpdateMemberProfile(ctx context.Context, exec bun.IDB, m *member.Member) (*member.Member, error) {
	p := m.Profile()
	idNum := ""
	if p.IDNumber() != nil {
		idNum = p.IDNumber().ToString()
	}
	mpID, _ := uuid.NewV7()
	mem := &models.MemberProfile{
		MemberProfileID: mpID,
		MemberID:        m.ID().Value(),
		MemberIDNumber:  idNum,
		DisplayName:     p.DisplayName().ToString(),
		Bio:             p.Bio().ToString(),
	}
	if _, err := exec.NewDelete().Model(&models.MemberLatestProfile{}).Where("member_id = ?", m.ID().Value()).
		Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := exec.NewInsert().Model(mem).Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := exec.NewInsert().Model(&models.MemberLatestProfile{
		MemberProfileID: mpID,
		MemberID:        m.ID().Value(),
	}).Exec(ctx); err != nil {
		return nil, err
	}
	return m, nil
}

func (d *driver) UpdateProfilePhoto(ctx context.Context, exec bun.IDB, aID account.ID, photo *media.UploadPhoto) error {
	peID, err := id.GenerateUUID()
	if err != nil {
		return err
	}
	if _, err = exec.NewDelete().Model(&models.AccountLatestPhotoEvent{}).Where("account_id = ?", aID.ToString()).
		Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.AccountPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
		EventType:           "upload",
	}).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.AccountLatestPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.AccountPhoto{
		AccountPhotoEventID: peID.Value(),
		PhotoID:             photo.ID().Value(),
		HostingTo:           photo.HostingTo().String(),
	}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (d *driver) RemoveProfilePhoto(ctx context.Context, exec bun.IDB, aID account.ID) error {
	peID, err := id.GenerateUUID()
	if err != nil {
		return err
	}
	if _, err = exec.NewDelete().Model(&models.AccountLatestPhotoEvent{}).Where("account_id = ?", aID.ToString()).
		Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.AccountPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
		EventType:           "remove",
	}).Exec(ctx); err != nil {
		return err
	}
	if _, err = exec.NewInsert().Model(&models.AccountLatestPhotoEvent{
		AccountPhotoEventID: peID.Value(),
		AccountID:           aID.Value(),
	}).Exec(ctx); err != nil {
		return err
	}
	return err
}
