package user

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(u *user.User) openapi.User
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(u *user.User) openapi.User {
	return openapi.User{
		UserId: u.AccountID().Value(),
		Email:  u.Email().ToString(),
		Name: openapi.OptString{
			Set:   u.HasName(),
			Value: u.Name().ToString(),
		},
		PhoneNumber: openapi.OptString{
			Set:   u.HasPhoneNumber(),
			Value: u.PhoneNumber().ToString(),
		},
	}
}
