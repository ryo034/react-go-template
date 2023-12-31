package request

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	accountPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/account/v1"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	userPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/user/v1"
)

type Adapter interface {
	Adapt(s *userPb.User) (*user.User, error)
	AdaptGender(g accountPb.Gender) (account.Gender, error)
}

func NewAdapter() Adapter {
	return &adapter{}
}

type adapter struct {
}

func (a *adapter) Adapt(u *userPb.User) (*user.User, error) {
	aID, err := account.NewID(u.UserId)
	if err != nil {
		return nil, err
	}
	email, err := account.NewEmail(u.Email)
	if err != nil {
		return nil, err
	}
	var phoneNumber *phone.Number = nil
	if u.Phone != nil {
		tmpPh, err := phone.NewPhoneNumber("")
		if err != nil {
			return nil, err
		}
		phoneNumber = &tmpPh
	}
	fn, err := account.NewFirstName(u.FirstName)
	if err != nil {
		return nil, err
	}
	ln, err := account.NewLastName(u.LastName)
	if err != nil {
		return nil, err
	}
	return user.NewUser(aID, email, phoneNumber, fn, ln), nil
}

func (a *adapter) AdaptGender(g accountPb.Gender) (account.Gender, error) {
	switch g {
	case accountPb.Gender_GENDER_MAN:
		return account.GenderMan, nil
	case accountPb.Gender_GENDER_WOMAN:
		return account.GenderWoman, nil
	case accountPb.Gender_GENDER_UNKNOWN:
		return account.GenderUnknown, nil
	}
	return account.GenderUnknown, fmt.Errorf("invalid gender: %v", g)
}
