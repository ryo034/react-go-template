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
	var na = openapi.OptString{Set: false}
	if u.HasName() {
		na.Set = true
		na.Value = u.Name().ToString()
	}
	var ph = openapi.OptString{Set: false}
	if u.HasPhoneNumber() {
		ph.Set = true
		ph.Value = u.PhoneNumber().ToNational()
	}
	return openapi.User{
		UserId:      u.AccountID().Value(),
		Email:       u.Email().ToString(),
		Name:        na,
		PhoneNumber: ph,
	}
}
