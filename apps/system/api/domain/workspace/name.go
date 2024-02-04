package workspace

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type Name struct {
	v string
}

const (
	InvalidWorkspaceName domainError.MessageKey = "invalid.workspace.name"
)

func NewName(v string) (Name, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidWorkspaceName, v)
	}
	if errs.IsNotEmpty() {
		return Name{}, errs
	}
	return Name{v}, nil
}

func (n Name) ToString() string {
	return n.v
}