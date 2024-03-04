package user

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
)

type User struct {
	accountId   account.ID
	email       account.Email
	name        *account.Name
	phoneNumber *phone.Number
	photo       *Photo
}

func NewUser(
	accountId account.ID,
	email account.Email,
	name *account.Name,
	phoneNumber *phone.Number,
	photo *Photo,
) *User {
	return &User{
		accountId,
		email,
		name,
		phoneNumber,
		photo,
	}
}
func NewTmpUser(accountId account.ID, email account.Email) *User {
	return &User{accountId, email, nil, nil, nil}
}

func (u *User) AccountID() account.ID {
	return u.accountId
}

func (u *User) Email() account.Email {
	return u.email
}

func (u *User) Name() *account.Name {
	return u.name
}

func (u *User) HasName() bool {
	return u.name != nil
}

func (u *User) HasNotName() bool {
	return u.name == nil
}

func (u *User) PhoneNumber() *phone.Number {
	return u.phoneNumber
}

func (u *User) HasPhoneNumber() bool {
	return u.phoneNumber != nil
}

func (u *User) HasNotPhoneNumber() bool {
	return u.phoneNumber == nil
}

func (u *User) UpdateName(name account.Name) *User {
	u.name = &name
	return u
}

func (u *User) Photo() *Photo {
	return u.photo
}

func (u *User) HasPhoto() bool {
	return u.photo != nil
}
