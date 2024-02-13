package time

import (
	"fmt"
	"log"
	"time"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

const (
	InvalidTime domainError.MessageKey = "invalid.time.time"
)

type Time struct {
	time.Time
}

func NewTime(t time.Time) Time {
	return Time{t}
}

func NewTimeAsLocal(t time.Time) Time {
	tz := datetime.CountryTz["Tokyo"]
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Println(err)
		log.Printf("LoadLocation Error: %s", tz)
		panic(err)
	}
	return Time{t.In(loc)}
}

func NewTimeFromWithSecondString(st string) (Time, error) {
	errs := validation.NewErrors()
	ct, err := time.Parse("15:04:05", st)
	if err != nil {
		errs.Append(InvalidTime, nil, st)
	}
	if errs.IsNotEmpty() {
		return Time{}, errs
	}
	return Time{ct}, nil
}

func (t Time) ToTime() time.Time {
	return t.Time
}

func (t Time) ToTimeString() string {
	return fmt.Sprintf("%d:%d", t.Time.Hour(), t.Time.Minute())
}

func (t Time) ToLocalTime() Time {
	loc, err := time.LoadLocation(datetime.CountryTz["Tokyo"])
	if err != nil {
		log.Println(err)
		log.Printf("ToLocalTime Error: %s", datetime.CountryTz["Tokyo"])
		panic(err)
	}
	return NewTime(t.In(loc))
}
