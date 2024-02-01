package workspace

import (
	"fmt"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"regexp"
)

type Subdomain struct {
	v string
}

const (
	InvalidWorkspaceSubdomain domainError.MessageKey = "invalid.workspace.subdomain"
)

const maxSubdomainLength = 63
const minSubdomainLength = 1

func isValidSubdomain(subdomain string) (bool, error) {
	if len(subdomain) < minSubdomainLength || len(subdomain) > maxSubdomainLength {
		return false, nil
	}
	return regexp.MatchString(`^[a-z0-9]+(-[a-z0-9]+)*$`, subdomain)
}

func NewSubdomain(v string) (Subdomain, error) {
	errs := validation.NewErrors()
	ok, err := isValidSubdomain(v)
	if err != nil || !ok {
		errs.Append(InvalidWorkspaceSubdomain, fmt.Sprintf("invalid workspace domain: %s", v))
		return Subdomain{}, errs
	}
	return Subdomain{v}, nil
}

func (d Subdomain) ToString() string {
	return d.v
}
