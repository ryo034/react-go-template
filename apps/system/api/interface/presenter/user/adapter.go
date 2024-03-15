package user

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(u *user.User) openapi.User
	AdaptForLeft(u *user.User) openapi.User
}

type adapter struct {
	sh storage.Handler
}

func NewAdapter(sh storage.Handler) Adapter {
	return &adapter{sh}
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

	var photo = openapi.OptURI{Set: false}
	if u.HasPhoto() {
		photo.Set = true
		if u.Photo().IsFirebase() {
			photo.Value = *u.Photo().URL()
		} else {
			ur, _ := a.sh.CreateAvatarPath(
				u.Photo().HostingTo(),
				u.AccountID(),
				u.Photo().ID())
			photo.Value = ur
		}
	}

	return openapi.User{
		UserId:      u.AccountID().Value(),
		Email:       u.Email().ToString(),
		Name:        na,
		PhoneNumber: ph,
		Photo:       photo,
	}
}

// AdaptForLeft is a method to adapt user for left user
// show only email
// mask name, phone number, photo
func (a *adapter) AdaptForLeft(u *user.User) openapi.User {
	var na = openapi.OptString{Set: true, Value: "Removed User"}
	var ph = openapi.OptString{Set: false}
	var photo = openapi.OptURI{Set: false}
	return openapi.User{
		UserId:      u.AccountID().Value(),
		Email:       u.Email().ToString(),
		Name:        na,
		PhoneNumber: ph,
		Photo:       photo,
	}
}
