package address

import (
	domainError "github.com/ryo034/react-go-template/packages/go/domain/shared/error"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/validation"
)

type Building struct {
	id    ID
	value string
}

const (
	InvalidAddressBuilding domainError.MessageKey = "invalid.address.building"
)

func NewBuilding(id ID, v string) (Building, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAddressBuilding, v)
	}
	if errs.IsNotEmpty() {
		return Building{}, errs
	}
	return Building{id, v}, nil
}

func (b Building) ToString() string {
	return b.value
}

func (b Building) ID() ID {
	return b.id
}
