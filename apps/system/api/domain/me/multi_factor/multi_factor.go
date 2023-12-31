package multi_factor

import (
	"github.com/ryo034/react-go-template/packages/go/domain/shared/datetime"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/phone"
)

type MultiFactor struct {
	phoneNumber         phone.Number
	enrollmentTimestamp datetime.Datetime
}

func NewMultiFactor(phoneNumber phone.Number, enrollmentTimestamp datetime.Datetime) MultiFactor {
	return MultiFactor{phoneNumber, enrollmentTimestamp}
}

func (m *MultiFactor) PhoneNumber() phone.Number {
	return m.phoneNumber
}

func (m *MultiFactor) EnrollmentTimestamp() datetime.Datetime {
	return m.enrollmentTimestamp
}
