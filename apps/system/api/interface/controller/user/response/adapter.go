package response

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	accountPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/account/v1"
	userPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/user/v1"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/account"
)

type Adapter interface {
	Adapt(u *user.User) *userPb.User
	AdaptGender(g account.Gender) accountPb.Gender
}

func NewAdapter() Adapter {
	return &adapter{}
}

type adapter struct {
}

func (a *adapter) AdaptGender(g account.Gender) accountPb.Gender {
	switch g {
	case account.GenderMan:
		return accountPb.Gender_GENDER_MAN
	case account.GenderWoman:
		return accountPb.Gender_GENDER_WOMAN
	case account.GenderUnknown:
		return accountPb.Gender_GENDER_UNKNOWN
	}
	return accountPb.Gender_GENDER_UNSPECIFIED
}

func (a *adapter) Adapt(u *user.User) *userPb.User {
	ph := ""
	if u.HasPhoneNumber() {
		ph = u.PhoneNumber().ToString()
	}
	return &userPb.User{
		UserId:    u.AccountID().ToString(),
		FirstName: u.FirstName().ToString(),
		LastName:  u.LastName().ToString(),
		Email:     u.Email().ToString(),
		Phone:     &ph,
	}
}
