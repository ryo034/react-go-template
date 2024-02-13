package id

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type StringID struct {
	v string
}

const InvalidStringID domainError.MessageKey = "invalid.string_id"

func NewStringID(v string) (StringID, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidStringID, nil, v)
	}
	if errs.IsNotEmpty() {
		return StringID{}, errs
	}
	return StringID{v}, nil
}

func (s StringID) ToString() string {
	return s.v
}
