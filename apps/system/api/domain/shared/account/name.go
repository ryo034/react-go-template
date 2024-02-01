package account

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type Name struct {
	v string
}

const (
	InvalidAccountName domainError.MessageKey = "invalid.account.name"
)

func NewName(v string) (Name, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAccountName, v)
	}
	if errs.IsNotEmpty() {
		return Name{}, errs
	}
	return Name{v}, nil
}

func (n Name) ToString() string {
	return n.v
}
