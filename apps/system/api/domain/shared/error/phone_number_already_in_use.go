package error

import "fmt"

type PhoneNumberAlreadyInUse struct {
	email string
}

func NewPhoneNumberAlreadyInUse(email string) *PhoneNumberAlreadyInUse {
	return &PhoneNumberAlreadyInUse{email: email}
}

func (p *PhoneNumberAlreadyInUse) PhoneNumber() string {
	return p.email
}

func (p *PhoneNumberAlreadyInUse) Error() string {
	return fmt.Sprintf("phone_number:%s is already in use", p.PhoneNumber())
}

func (p *PhoneNumberAlreadyInUse) MessageKey() MessageKey {
	return PhoneNumberAlreadyInUseMessageKey
}

func (p *PhoneNumberAlreadyInUse) Code() string {
	return "400-" + string(PhoneNumberInUseCodeKey)
}
