package address

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type City struct {
	id    ID
	value string
}

const (
	InvalidAddressCity domainError.MessageKey = "invalid.address.city"
)

func NewCity(id ID, v string) (City, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAddressCity, v)
	}
	if errs.IsNotEmpty() {
		return City{}, errs
	}
	return City{id, v}, nil
}

func (c City) ToString() string {
	return c.value
}

func (c City) ID() ID {
	return c.id
}
