package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	AdaptTmp(sa *models.SystemAccount) (*user.User, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) AdaptTmp(sa *models.SystemAccount) (*user.User, error) {
	em, err := account.NewEmail(sa.Emails[0].Email)
	if err != nil {
		return nil, err
	}
	aID := account.NewIDFromUUID(sa.SystemAccountID)

	var an *account.Name = nil
	if sa.Profile != nil && sa.Profile.Name != "" {
		tmpAn, err := account.NewName(sa.Profile.Name)
		if err != nil {
			return nil, err
		}
		an = &tmpAn
	}

	var pn *phone.Number = nil
	if sa.PhoneNumbers != nil {
		tmpPn, err := phone.NewInternationalPhoneNumber(sa.PhoneNumbers[0].PhoneNumber, sa.PhoneNumbers[0].CountryCode)
		if err != nil {
			return nil, err
		}
		pn = &tmpPn
	}
	return user.NewUser(aID, em, an, pn), nil
}
