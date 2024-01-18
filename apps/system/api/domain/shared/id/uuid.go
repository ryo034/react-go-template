package id

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type UUID struct {
	v uuid.UUID
}

func NewUUIDFromString(v string) (UUID, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidStringID, v)
	}
	if uuid.Validate(v) != nil {
		errs.Append(InvalidStringID, v)
	}
	if errs.IsNotEmpty() {
		return UUID{}, errs
	}
	return UUID{uuid.MustParse(v)}, nil
}

func NewUUID(v uuid.UUID) UUID {
	return UUID{v}
}

func (u UUID) ToString() string {
	return u.v.String()
}

func (u UUID) Value() uuid.UUID {
	return u.v
}
