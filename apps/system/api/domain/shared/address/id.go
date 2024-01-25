package address

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/id"

type ID struct {
	id.UUID
}

func NewID(v string) (ID, error) {
	i, err := id.NewUUIDFromString(v)
	if err != nil {
		return ID{}, err
	}
	return ID{i}, nil
}
