package member

import (
	"context"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"github.com/uptrace/bun"
)

type Driver interface {
	Find(ctx context.Context, exec bun.IDB, mID member.ID) (*models.Member, error)
}

type driver struct {
}

func NewDriver() Driver {
	return &driver{}
}

func (p *driver) Find(ctx context.Context, exec bun.IDB, mID member.ID) (*models.Member, error) {
	m := models.Member{}
	err := exec.
		NewSelect().
		Model(&m).
		Relation("Profile").
		Relation("Role").
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
	return &m, nil
}
