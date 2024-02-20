package me

import (
	"context"
	"fmt"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
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
	UpdateMember(ctx context.Context, exec bun.IDB, m *me.Me) error
	UpdateProfile(ctx context.Context, exec bun.IDB, usr *user.User) error
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
		Relation("Workspace").
		Relation("Workspace.Detail").
		Relation("SystemAccount").
		Relation("SystemAccount.Profile").
		Relation("SystemAccount.PhoneNumber").
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
	_, err = exec.
		NewInsert().
		Model(m).
		Exec(ctx)
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
		Relation("Profile").
		Relation("PhoneNumber").
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
		Relation("Profile").
		Relation("PhoneNumber").
		Where("sa.system_account_id = ?", aID.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc, nil
}

func (d *driver) FindByEmail(ctx context.Context, exec bun.IDB, email account.Email) (*models.SystemAccount, error) {
	sysAcc := &models.SystemAccount{}
	err := exec.
		NewSelect().
		Model(sysAcc).
		Relation("Profile").
		Relation("PhoneNumber").
		Where("email = ?", email.ToString()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return sysAcc, nil
}

func (d *driver) UpdateMember(ctx context.Context, exec bun.IDB, m *me.Me) error {
	mem := &models.MemberProfile{
		MemberID:       m.Member().ID().Value(),
		MemberIDNumber: m.Member().IDNumber().ToString(),
		DisplayName:    m.Member().DisplayName().ToString(),
	}
	_, err := exec.
		NewUpdate().
		Model(mem).
		WherePK().
		Exec(ctx)
	return err
}

func (d *driver) UpdateProfile(ctx context.Context, exec bun.IDB, usr *user.User) error {
	mem := &models.SystemAccountProfile{
		SystemAccountID: usr.AccountID().Value(),
		Name:            usr.Name().ToString(),
		Email:           usr.Email().ToString(),
	}
	res, err := exec.
		NewUpdate().
		Model(mem).
		WherePK().
		Exec(ctx)
	fmt.Println(res)
	return err
}
