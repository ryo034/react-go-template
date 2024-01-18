package request

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
)

type Adapter interface {
	Adapt(s *openapi.User) (*user.User, error)
}

func NewAdapter() Adapter {
	return &adapter{}
}

type adapter struct {
}

func (a *adapter) Adapt(u *openapi.User) (*user.User, error) {
	aID := account.NewIDFromUUID(u.UserId)
	email, err := account.NewEmail(u.Email)
	if err != nil {
		return nil, err
	}
	var phoneNumber *phone.Number = nil

	n, err := account.NewName(u.Name)
	if err != nil {
		return nil, err
	}
	return user.NewUser(aID, email, n, phoneNumber), nil
}
