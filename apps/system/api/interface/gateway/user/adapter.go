package user

import (
	"net/url"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	AdaptTmp(u *models.Account) (*user.User, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) AdaptTmp(u *models.Account) (*user.User, error) {
	aID, err := account.NewID(u.AccountID.String())
	if err != nil {
		return nil, err
	}
	email, err := account.NewEmail(u.Email.AccountEmail.Email)
	if err != nil {
		return nil, err
	}

	var nm *account.Name = nil
	if u.Name != nil && u.Name.AccountName != nil {
		name, err := account.NewName(u.Name.AccountName.Name)
		if err != nil {
			return nil, err
		}
		nm = &name
	}

	var pn *phone.Number = nil
	if u.PhoneNumber != nil && u.PhoneNumber.AccountPhoneNumber != nil {
		tmpPn, err := phone.NewInternationalPhoneNumber(u.PhoneNumber.AccountPhoneNumber.PhoneNumber, u.PhoneNumber.AccountPhoneNumber.CountryCode)
		if err != nil {
			return nil, err
		}
		pn = &tmpPn
	}

	var pho *user.Photo = nil
	// 最新の登録されているプロバイダ情報から写真URLを取得
	if u.PhotoEvent == nil && u.AuthProviders != nil && u.AuthProviders[0].PhotoURL != "" {
		phoURL, err := url.Parse(u.AuthProviders[0].PhotoURL)
		if err != nil {
			return nil, err
		}
		pho = user.NewPhotoFromFirebase(phoURL)
	}
	if u.PhotoEvent != nil && u.PhotoEvent.AccountPhotoEvent.EventType == "upload" {
		pho = user.NewPhoto(
			media.NewIDFromUUID(u.PhotoEvent.AccountPhotoEvent.Photo.PhotoID),
			media.HostingToR2,
			nil,
		)
	}

	return user.NewUser(aID, email, nm, pn, pho), nil
}
