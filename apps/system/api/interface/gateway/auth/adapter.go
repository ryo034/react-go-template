package auth

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	AdaptTmp(sa *models.Account) (*user.User, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) AdaptTmp(sa *models.Account) (*user.User, error) {
	em, err := account.NewEmail(sa.Email.AccountEmail.Email)
	if err != nil {
		return nil, err
	}
	aID := account.NewIDFromUUID(sa.AccountID)

	var an *account.Name = nil
	if sa.Name != nil && sa.Name.AccountName.Name != "" {
		tmpAn, err := account.NewName(sa.Name.AccountName.Name)
		if err != nil {
			return nil, err
		}
		an = &tmpAn
	}

	var pn *phone.Number = nil
	if sa.PhoneNumber != nil {
		tmpPn, err := phone.NewInternationalPhoneNumber(sa.PhoneNumber.AccountPhoneNumber.PhoneNumber, sa.PhoneNumber.AccountPhoneNumber.CountryCode)
		if err != nil {
			return nil, err
		}
		pn = &tmpPn
	}

	var pho *user.Photo = nil
	if sa.PhotoEvent != nil && sa.PhotoEvent.AccountPhotoEvent.EventType == "upload" {
		tmpPho, err := user.NewPhotoFromString(sa.PhotoEvent.AccountPhotoEvent.Photo.PhotoPath)
		if err != nil {
			return nil, err
		}
		pho = &tmpPho
	}
	return user.NewUser(aID, em, an, pn, pho), nil
}
