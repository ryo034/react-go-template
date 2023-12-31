package address

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type Country struct {
	code  ID
	value string
}

const (
	InvalidAddressCountry domainError.MessageKey = "invalid.address.country"
)

func NewCountry(code ID, v string) (Country, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAddressCountry, v)
	}
	if errs.IsNotEmpty() {
		return Country{}, errs
	}
	return Country{code: code, value: v}, nil
}

func (c Country) ToString() string {
	return c.value
}

func (c Country) Code() ID {
	return c.code
}
