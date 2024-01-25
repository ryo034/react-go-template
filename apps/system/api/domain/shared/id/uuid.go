package id

import (
	"github.com/google/uuid"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

type UUID struct {
	v uuid.UUID
}

const InvalidUUID domainError.MessageKey = "invalid.uuid"

func NewUUIDFromString(v string) (UUID, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidUUID, v)
	}
	if uuid.Validate(v) != nil {
		errs.Append(InvalidUUID, v)
	}
	if errs.IsNotEmpty() {
		return UUID{}, errs
	}
	return UUID{uuid.MustParse(v)}, nil
}

func NewUUID(v uuid.UUID) UUID {
	return UUID{v}
}

func GenerateUUID() (UUID, error) {
	errs := validation.NewErrors()
	i, err := uuid.NewV7()
	if err != nil {
		errs.Append(InvalidUUID, i)
	}
	if errs.IsNotEmpty() {
		return UUID{}, errs
	}
	return UUID{i}, nil
}

func (u UUID) ToString() string {
	return u.v.String()
}

func (u UUID) Value() uuid.UUID {
	return u.v
}
