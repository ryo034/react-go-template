package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	"time"
)

type VerifiedAt struct {
	v datetime.Datetime
}

func NewVerifiedAt(v datetime.Datetime) VerifiedAt {
	return VerifiedAt{v}
}

func (e VerifiedAt) Value() datetime.Datetime {
	return e.v
}

func (e VerifiedAt) IsVerified() bool {
	return e.v.IsBefore(time.Now())
}
