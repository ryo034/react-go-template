package phone

import (
	"github.com/ttacon/libphonenumber"
	"regexp"
	"strings"

	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

const (
	InvalidPhoneNumber domainError.MessageKey = "invalid.phone_number"
)

type Number struct {
	value string
}

const phoneNumberRegex = `^0[789]0\d{8}$`

func NewPhoneNumber(v string) (Number, error) {
	errs := validation.NewErrors()
	if !regexp.MustCompile(phoneNumberRegex).MatchString(v) {
		errs.Append(InvalidPhoneNumber, v)
	}
	if errs.IsNotEmpty() {
		return Number{}, errs
	}
	return Number{v}, nil
}

func NewInternationalPhoneNumber(v string) (Number, error) {
	errs := validation.NewErrors()
	num, err := libphonenumber.Parse(v, "JP")
	if err != nil {
		return Number{}, err
	}
	formatted := libphonenumber.Format(num, libphonenumber.NATIONAL)
	noHyphen := strings.ReplaceAll(formatted, "-", "")
	ph, err := NewPhoneNumber(noHyphen)
	if err != nil {
		errs.Append(InvalidPhoneNumber, v)
	}
	if errs.IsNotEmpty() {
		return Number{}, errs
	}
	return ph, nil
}

func (e Number) ToString() string {
	return e.value
}

func (e Number) ToInternationalNumberString() (string, error) {
	num, err := libphonenumber.Parse(e.ToString(), "JP")
	if err != nil {
		return "", err
	}
	formatted := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
	noHyphen := strings.ReplaceAll(formatted, "-", "")
	return strings.ReplaceAll(noHyphen, " ", ""), nil
}
