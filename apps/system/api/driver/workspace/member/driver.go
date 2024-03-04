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
	return &m, nil
}
