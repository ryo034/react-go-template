package media

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type Path string

const (
	InvalidMediaPath domainError.MessageKey = "invalid.media.path"
)

func NewPath(v string) (Path, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidMediaPath, v)
	}
	if v[0] == '/' {
		errs.Append(InvalidMediaPath, v)
	}
	if errs.IsNotEmpty() {
		return "", errs
	}
	return Path(v), nil
}

func (p Path) String() string {
	return string(p)
}
