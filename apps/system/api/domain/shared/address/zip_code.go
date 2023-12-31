package address

import (
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

const (
	InvalidAddressZipCode domainError.MessageKey = "invalid.address.zip_code"
)

type ZipCode struct {
	value string
}

func NewZipCode(v string) (ZipCode, error) {
	errs := validation.NewErrors()
	if v == "" {
		errs.Append(InvalidAddressZipCode, v)
	}
	if errs.IsNotEmpty() {
		return ZipCode{}, errs
	}
	return ZipCode{v}, nil
}

func (z ZipCode) ToString() string {
	return z.value
}
