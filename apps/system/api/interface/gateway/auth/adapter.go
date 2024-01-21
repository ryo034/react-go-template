package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	Adapt(sa *models.SystemAccount) (*user.User, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(sa *models.SystemAccount) (*user.User, error) {
	em, err := account.NewEmail(sa.Profile.Email)
	if err != nil {
		return nil, err
	}
	return user.NewTmpUser(account.NewIDFromUUID(sa.SystemAccountID), em), nil
}
