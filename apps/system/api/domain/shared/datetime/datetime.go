package datetime

import (
	"log"
	"time"

	domainError "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
)

const (
	InvalidDate     domainError.MessageKey = "invalid.datetime.date"
	InvalidDatetime domainError.MessageKey = "invalid.datetime.datetime"
)

type Datetime struct {
	value time.Time
}

var CountryTz = map[string]string{
	"Hungary": "Europe/Budapest",
	"Egypt":   "Africa/Cairo",
	"Tokyo":   "Asia/Tokyo",
}

func NewDatetime(t time.Time) Datetime {
	return Datetime{t}
}

func Now() Datetime {
	return Datetime{time.Now()}
}

func NewDatetimeFromString(st string) (Datetime, error) {
	errs := validation.NewErrors()
	ct, err := time.Parse("2006-01-02 15:04:05", st)
	if err != nil {
		errs.Append(InvalidDatetime, st)
	}
	if errs.IsNotEmpty() {
		return Datetime{}, errs
	}
	return Datetime{ct}, nil
}

func NewDateFromString(st string) (Datetime, error) {
	errs := validation.NewErrors()
	ct, err := time.Parse("2006-01-02", st)
	if err != nil {
		errs.Append(InvalidDate, st)
	}
	if errs.IsNotEmpty() {
		return Datetime{}, errs
	}
	return Datetime{ct}, nil
}

func NewDateFromInt64(t int64) (Datetime, error) {
	return Datetime{time.Unix(t, 0)}, nil
}

func NewDateRFC3339FromString(st string) (Datetime, error) {
	errs := validation.NewErrors()
	ct, err := time.Parse(time.RFC3339, st)
	if err != nil {
		errs.Append(InvalidDate, st)
	}
	if errs.IsNotEmpty() {
		return Datetime{}, errs
	}
	return Datetime{ct}, nil
}

func (d Datetime) IsAfter(t time.Time) bool {
	return d.value.After(t)
}

func (d Datetime) IsBefore(t time.Time) bool {
	return d.value.Before(t)
}

func (d Datetime) ToTime() time.Time {
	return d.value
}

func (d Datetime) ToDate() time.Time {
	t := d.value
	t = t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour)
	return t
}

func (d Datetime) ToDateString() string {
	if d.value.IsZero() {
		return ""
	}
	return d.ToDate().Format("2006-01-02")
}

func (d Datetime) ToDatetimeString() string {
	if d.value.IsZero() {
		return ""
	}
	return d.ToDate().Format("2006-01-02 15:04:05")
}

func (d Datetime) ToLocalDate() time.Time {
	loc, err := time.LoadLocation(CountryTz["Tokyo"])
	if err != nil {
		log.Println(err)
		log.Printf("ToLocalTime Error: %s", CountryTz["Tokyo"])
		panic(err)
	}
	t := d.value
	return t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour).In(loc)
}

func (d Datetime) ToLocalTime() time.Time {
	tz := CountryTz["Tokyo"]
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Println(err)
		log.Printf("LoadLocation Error: %s", tz)
		panic(err)
	}
	return d.value.In(loc)
}
