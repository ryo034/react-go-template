package account

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type LastName string

const (
	InvalidAccountLastName domainError.MessageKey = "invalid.account.last_name"
)

func NewLastName(v string) (LastName, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAccountLastName, v)
	}
	if errs.IsNotEmpty() {
		return "", errs
	}
	return LastName(v), nil
}

func (v LastName) ToString() string {
	return string(v)
}
