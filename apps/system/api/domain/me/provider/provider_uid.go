package provider

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/id"
)

type UID struct {
	v id.StringID
}

func NewUID(v string) (UID, error) {
	nid, err := id.NewStringID(v)
	if err != nil {
		return UID{}, err
	}
	return UID{nid}, nil
}

func (u UID) ToString() string {
	return u.v.ToString()
}
