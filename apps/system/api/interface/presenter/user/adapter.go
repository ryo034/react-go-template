package user

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type Adapter interface {
	Adapt(u *user.User) openapi.User
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
