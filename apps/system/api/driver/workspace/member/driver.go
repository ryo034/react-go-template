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
		Relation("SystemAccount").
		Relation("SystemAccount.Profile").
		Relation("SystemAccount.PhoneNumber").
		Where("m.member_id = ?", mID.Value()).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
