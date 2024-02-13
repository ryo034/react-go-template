package address

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type Prefecture struct {
	id    ID
	value string
}

const (
	InvalidAddressPrefecture domainError.MessageKey = "invalid.address.prefecture"
)

func NewPrefecture(id ID, v string) (Prefecture, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAddressPrefecture, nil, v)
	}
	if errs.IsNotEmpty() {
		return Prefecture{}, errs
	}
	return Prefecture{id, v}, nil
}

func (p Prefecture) ToString() string {
	return p.value
}

func (p Prefecture) ID() ID {
	return p.id
}
