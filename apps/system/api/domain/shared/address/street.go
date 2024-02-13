package address

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type Street struct {
	id    ID
	value string
}

const (
	InvalidAddressStreet domainError.MessageKey = "invalid.address.street"
)

func NewStreet(id ID, v string) (Street, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAddressStreet, nil, v)
	}
	if errs.IsNotEmpty() {
		return Street{}, errs
	}
	return Street{id, v}, nil
}

func (s Street) ToString() string {
	return s.value
}

func (s Street) ID() ID {
	return s.id
}
