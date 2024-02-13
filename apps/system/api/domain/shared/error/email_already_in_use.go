package error

import "fmt"

type EmailAlreadyInUse struct {
	email string
}

func NewEmailAlreadyInUse(email string) *EmailAlreadyInUse {
	return &EmailAlreadyInUse{email: email}
}

func (e *EmailAlreadyInUse) Email() string {
	return e.email
}

func (e *EmailAlreadyInUse) Error() string {
	return fmt.Sprintf("email:%s is already in use", e.Email())
}

func (e *EmailAlreadyInUse) MessageKey() MessageKey {
	return EmailAlreadyInUseMessageKey
}

func (e *EmailAlreadyInUse) Code() string {
	return "400-" + string(EmailAlreadyInUseCodeKey)
}
