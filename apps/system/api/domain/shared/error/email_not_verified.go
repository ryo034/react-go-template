package error

import "fmt"

type EmailNotVerified struct {
	email string
}

func NewEmailNotVerified(email string) *EmailNotVerified {
	return &EmailNotVerified{email: email}
}

func (e *EmailNotVerified) Email() string {
	return e.email
}

func (e *EmailNotVerified) Error() string {
	return fmt.Sprintf("email:%s not verified", e.Email())
}

func (e *EmailNotVerified) MessageKey() MessageKey {
	return EmailNotVerifiedMessageKey
}
