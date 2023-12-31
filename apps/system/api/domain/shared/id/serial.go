package id

import (
	"github.com/spf13/cast"
)

type SerialID struct {
	id uint
}

func NewSerialID(id uint) SerialID {
	return SerialID{id}
}

func (s SerialID) ToString() string {
	return cast.ToString(s.id)
}

func (s SerialID) ToUint() uint {
	return s.id
}

func (s SerialID) ToUint32() uint32 {
	return cast.ToUint32(s.id)
}

func (s SerialID) ToUint8() uint8 {
	return cast.ToUint8(s.id)
}

func (s SerialID) ToInt8() int8 {
	return cast.ToInt8(s.id)
}

func (s SerialID) ToUint64() uint64 {
	return cast.ToUint64(s.id)
}
