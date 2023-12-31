package address

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/id"

type ID struct {
	id.SortableID
}

func NewID(v string) (ID, error) {
	i, err := id.NewSortableID(v)
	if err != nil {
		return ID{}, err
	}
	return ID{i}, nil
}

func GenID() ID {
	return ID{id.GenStringID()}
}
