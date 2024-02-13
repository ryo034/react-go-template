package error

import "fmt"

type InvalidEmail struct {
	email string
}

func NewInvalidEmail(email string) *InvalidEmail {
	return &InvalidEmail{email: email}
}

func (e *InvalidEmail) Email() string {
	return e.email
}

func (e *InvalidEmail) Error() string {
	return fmt.Sprintf("email:%s is already in use", e.Email())
}

func (e *InvalidEmail) MessageKey() MessageKey {
	return InvalidEmailMessageKey
}

func (e *InvalidEmail) Code() string {
	return "400-" + string(InvalidEmailCodeKey)
}
