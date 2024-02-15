package user

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	AdaptTmp(u *models.SystemAccount) (*user.User, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) AdaptTmp(u *models.SystemAccount) (*user.User, error) {
	aID, err := account.NewID(u.SystemAccountID.String())
	if err != nil {
		return nil, err
	}
	email, err := account.NewEmail(u.Profile.Email)
	if err != nil {
		return nil, err
	}
	var nm *account.Name = nil
	if u.Profile.Name != "" {
		name, err := account.NewName(u.Profile.Name)
		if err != nil {
			return nil, err
		}
		nm = &name
	}

	var pn *phone.Number = nil
	if u.PhoneNumber != nil {
		tmpPn, err := phone.NewPhoneNumber(u.PhoneNumber.PhoneNumber)
		if err != nil {
			return nil, err
		}
		pn = &tmpPn
	}
	return user.NewUser(aID, email, nm, pn), nil
}
