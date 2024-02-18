package invitation

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	"time"
)

const ExpiredDays = 7

type ExpiredAt struct {
	v datetime.Datetime
}

func NewExpiredAt(v datetime.Datetime) ExpiredAt {
	return ExpiredAt{v}
}

func (e ExpiredAt) Value() datetime.Datetime {
	return e.v
}

func (e ExpiredAt) IsExpired() bool {
	return e.v.IsBefore(time.Now())
}

func (e ExpiredAt) IsNotExpired() bool {
	return !e.IsExpired()
}

func GenerateExpiredAt() ExpiredAt {
	return ExpiredAt{datetime.NewDatetime(time.Now().AddDate(0, 0, ExpiredDays))}
}
