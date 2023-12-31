package account

import (
	"github.com/ryo034/react-go-template/packages/go/domain/shared/id"
)

// ID used by external SaaS such as Firebase
type ID struct {
	id.StringID
}

func NewID(v string) (ID, error) {
	i, err := id.NewStringID(v)
	if err != nil {
		return ID{}, err
	}
	return ID{i}, nil
}
