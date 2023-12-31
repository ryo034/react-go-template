package error

import (
	"fmt"
)

type InvalidAddress struct {
	address string
}

func NewInvalidAddress(address string) *InvalidAddress {
	return &InvalidAddress{address}
}

func (e *InvalidAddress) Address() string {
	return e.address
}

func (e *InvalidAddress) Error() string {
	return fmt.Sprintf("invalid address: %s", e.Address())
}

func (e *InvalidAddress) MessageKey() MessageKey {
	return InvalidAddressMessageKey
}
