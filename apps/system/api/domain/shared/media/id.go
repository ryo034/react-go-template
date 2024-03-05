package media

import (
	"github.com/google/uuid"
)

type ID struct {
	v uuid.UUID
}

func NewIDFromUUID(v uuid.UUID) ID {
	return ID{v}
}

func (i ID) Value() uuid.UUID {
	return i.v
}

func (i ID) String() string {
	return i.v.String()
}
