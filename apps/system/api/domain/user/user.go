package user

import (
	"github.com/ryo034/react-go-template/packages/go/domain/shared/account"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/phone"
)

type User struct {
	accountId   account.ID
	email       account.Email
	phoneNumber *phone.Number
	firstName   account.FirstName
	lastName    account.LastName
}

func NewUser(
	accountId account.ID,
	email account.Email,
	phoneNumber *phone.Number,
	firstName account.FirstName,
	lastName account.LastName,
) *User {
	return &User{
		accountId,
		email,
		phoneNumber,
		firstName,
		lastName,
	}
}

func (s *User) AccountID() account.ID {
	return s.accountId
}

func (s *User) Email() account.Email {
	return s.email
}

func (s *User) PhoneNumber() *phone.Number {
	return s.phoneNumber
}

func (s *User) HasPhoneNumber() bool {
	return s.phoneNumber != nil
}

func (s *User) Name() string {
	return s.LastName().ToString() + " " + s.FirstName().ToString()
}

func (s *User) FirstName() account.FirstName {
	return s.firstName
}

func (s *User) LastName() account.LastName {
	return s.lastName
}
