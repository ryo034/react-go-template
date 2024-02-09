package workspace

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"regexp"
	"strings"
)

type Name struct {
	v string
}

const (
	InvalidWorkspaceName domainError.MessageKey = "invalid.workspace.name"
	MaxLength                                   = 255
	Regex                                       = `^[a-zA-Z0-9ぁ-んァ-ヶー一-龠\-]+$`
)

func NewName(v string) (Name, error) {
	errs := validation.NewErrors()
	trimmed := strings.TrimSpace(v)
	if trimmed == "" {
		errs.Append(InvalidWorkspaceName, v)
	} else if len(trimmed) > MaxLength {
		errs.Append(InvalidWorkspaceName, v)
	} else {
		ok, err := regexp.MatchString(Regex, trimmed)
		if err != nil || !ok {
			errs.Append(InvalidWorkspaceName, v)
		}
	}
	if errs.IsNotEmpty() {
		return Name{}, errs
	}
	return Name{trimmed}, nil
}

func (n Name) ToString() string {
	return n.v
}
