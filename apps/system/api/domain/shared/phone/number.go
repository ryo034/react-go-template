package phone

import (
	"github.com/ttacon/libphonenumber"

	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

const (
	InvalidPhoneNumber domainError.MessageKey = "invalid.phone_number"
)

type Number struct {
	value  string
	region string
}

func isValidJapanesePhoneNumber(num *libphonenumber.PhoneNumber) bool {
	// 例: 020や050で始まる番号を無効とするロジック
	numberStr := libphonenumber.Format(num, libphonenumber.NATIONAL)
	prefix := numberStr[:3] // 日本の番号であれば、プレフィックスは3桁です
	return prefix != "020" && prefix != "050"
}

func NewInternationalPhoneNumber(v string, region string) (Number, error) {
	if region == "" {
		region = "JP"
	}
	errs := validation.NewErrors()
	num, err := libphonenumber.Parse(v, region)
	if err != nil || !libphonenumber.IsValidNumber(num) {
		errs.Append(InvalidPhoneNumber, v)
		return Number{}, errs
	}
	if !isValidJapanesePhoneNumber(num) {
		errs.Append(InvalidPhoneNumber, v)
		return Number{}, errs
	}
	return Number{libphonenumber.Format(num, libphonenumber.E164), region}, nil
}

func (n Number) ToE164() string {
	return n.value
}

func (n Number) ToNational() string {
	num, _ := libphonenumber.Parse(n.value, n.region)
	return libphonenumber.Format(num, libphonenumber.NATIONAL)
}

func (n Number) Region() string {
	return n.region
}
