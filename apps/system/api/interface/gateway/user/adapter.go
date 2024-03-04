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
	email, err := account.NewEmail(u.Email.SystemAccountEmail.Email)
	if err != nil {
		return nil, err
	}

	var nm *account.Name = nil
	if u.Name != nil {
		name, err := account.NewName(u.Name.SystemAccountName.Name)
		if err != nil {
			return nil, err
		}
		nm = &name
	}

	var pn *phone.Number = nil
	if u.PhoneNumber != nil {
		tmpPn, err := phone.NewInternationalPhoneNumber(u.PhoneNumber.SystemAccountPhoneNumber.PhoneNumber, u.PhoneNumber.SystemAccountPhoneNumber.CountryCode)
		if err != nil {
			return nil, err
		}
		pn = &tmpPn
	}

	var pho *user.Photo = nil
	if u.PhotoEvent != nil && u.PhotoEvent.SystemAccountPhotoEvent.EventType == "upload" {
		tmpPho, err := user.NewPhotoFromString(u.PhotoEvent.SystemAccountPhotoEvent.Photo.PhotoPath)
		if err != nil {
			return nil, err
		}
		pho = &tmpPho
	}

	return user.NewUser(aID, email, nm, pn, pho), nil
}
