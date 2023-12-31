package account

import (
	"regexp"

	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

const (
	InvalidEmail domainError.MessageKey = "invalid.email"
)

type Email struct {
	value string
}

const emailRegex = `^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`

func NewEmail(v string) (Email, error) {
	errs := validation.NewErrors()
	if !regexp.MustCompile(emailRegex).MatchString(v) {
		errs.Append(InvalidEmail, v)
	}
	if errs.IsNotEmpty() {
		return Email{}, errs
	}
	return Email{v}, nil
}

func (e Email) ToString() string {
	return e.value
}
