package workspace

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/id"
)

type ID struct {
	id.UUID
}

func NewIDFromUUID(v uuid.UUID) ID {
	return ID{id.NewUUID(v)}
}

func GenerateID() (ID, error) {
	i, err := id.GenerateUUID()
	if err != nil {
		return ID{}, err
	}
	return ID{i}, nil
}
