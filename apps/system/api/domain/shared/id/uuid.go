package id

import (
	"encoding/base32"
	"github.com/google/uuid"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"strings"
)

type UUID struct {
	v uuid.UUID
}

const InvalidUUID domainError.MessageKey = "invalid.uuid"

func NewUUIDFromString(v string) (UUID, error) {
	errs := validation.NewErrors()
	if uuid.Validate(v) != nil {
		errs.Append(InvalidUUID, nil, v)
	}
	if errs.IsNotEmpty() {
		return UUID{}, errs
	}
	return UUID{uuid.MustParse(v)}, nil
}

func NewFromFriendlyString(v string) (UUID, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidUUID, nil, v)
		return UUID{}, errs
	}
	uuid4Binary, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(strings.ToUpper(v))
	if err != nil {
		errs.Append(InvalidUUID, nil, v)
	}
	if errs.IsNotEmpty() {
		return UUID{}, errs
	}
	return UUID{uuid.Must(uuid.FromBytes(uuid4Binary))}, nil
}

func NewUUID(v uuid.UUID) UUID {
	return UUID{v}
}

func GenerateUUID() (UUID, error) {
	errs := validation.NewErrors()
	i, err := uuid.NewV7()
	if err != nil {
		errs.Append(InvalidUUID, nil, i)
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

// ToFriendlyString return base32 encoded string
func (u UUID) ToFriendlyString() string {
	uuid4Binary, _ := u.v.MarshalBinary()
	uuid4EncodedBase32 := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(uuid4Binary)
	return strings.ToLower(uuid4EncodedBase32)
}
