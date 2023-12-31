package account

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type FirstName string

const (
	InvalidAccountFirstName domainError.MessageKey = "invalid.account.first_name"
)

func NewFirstName(v string) (FirstName, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAccountFirstName, v)
	}
	if errs.IsNotEmpty() {
		return "", errs
	}
	return FirstName(v), nil
}

func (v FirstName) ToString() string {
	return string(v)
}
