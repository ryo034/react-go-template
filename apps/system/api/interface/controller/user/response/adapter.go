package response

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(u *user.User) *openapi.User
}

func NewAdapter() Adapter {
	return &adapter{}
}

type adapter struct {
}

func (a *adapter) Adapt(u *user.User) *openapi.User {
	ph := ""
	if u.HasPhoneNumber() {
		ph = u.PhoneNumber().ToString()
	}
	return &openapi.User{
		UserId:    u.AccountID().ToString(),
		Email:     u.Email().ToString(),
		FirstName: u.FirstName().ToString(),
		LastName:  u.LastName().ToString(),
		PhoneNumber: openapi.OptString{
			Value: ph,
			Set:   false,
		},
	}
}
