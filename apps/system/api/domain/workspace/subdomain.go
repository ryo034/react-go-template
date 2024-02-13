package workspace

import (
	"fmt"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"regexp"
	"strings"
)

type Subdomain struct {
	v string
}

const (
	InvalidWorkspaceSubdomain domainError.MessageKey = "invalid.workspace.subdomain"
)

const maxLength = 63
const minLength = 1

func isValidSubdomain(subdomain string) (bool, error) {
	if len(subdomain) < minLength || len(subdomain) > maxLength {
		return false, nil
	}
	return regexp.MatchString(`^[a-z0-9]+(-[a-z0-9]+)*$`, subdomain)
}

func NewSubdomain(v string) (Subdomain, error) {
	errs := validation.NewErrors()
	trimmed := strings.TrimSpace(v)
	ok, err := isValidSubdomain(trimmed)
	if err != nil || !ok {
		errs.Append(InvalidWorkspaceSubdomain, nil, fmt.Sprintf("invalid workspace domain: %s", v))
		return Subdomain{}, errs
	}
	return Subdomain{trimmed}, nil
}

func (d Subdomain) ToString() string {
	return d.v
}
