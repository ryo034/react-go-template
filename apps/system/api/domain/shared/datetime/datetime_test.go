package datetime

import (
	"reflect"
	"testing"
	"time"
)

var now = time.Now().UTC()

func Test_Datetime_ToDate_OK(t *testing.T) {
	tests := []struct {
		name   string
		fields time.Time
		want   time.Time
	}{
		{name: "Return Time", fields: time.Time{}, want: time.Time{}},
		{
			name:   "Midnight",
			fields: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
			want:   time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			name:   "Noon",
			fields: time.Date(2023, 4, 10, 12, 0, 0, 0, time.UTC),
			want:   time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			name:   "Evening",
			fields: time.Date(2023, 4, 10, 23, 59, 59, 0, time.UTC),
			want:   time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Datetime{
				value: tt.fields,
			}
			if got := d.ToDate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Datetime_ToDateString_OK(t *testing.T) {
	mockTime := time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name   string
		fields time.Time
		want   string
	}{
		{name: "if empty, return empty string", fields: time.Time{}, want: ""},
		{name: "Return Date String", fields: mockTime, want: "2020-05-01"},
		{
			name:   "Specific Date",
			fields: time.Date(2023, 4, 10, 15, 30, 0, 0, time.UTC),
			want:   "2023-04-10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Datetime{
				value: tt.fields,
			}
			if got := d.ToDateString(); got != tt.want {
				t.Errorf("ToDateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_ToDatetimeString_OK(t *testing.T) {
	mockTime := time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC)

	type fields struct {
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "if empty, return empty string", fields: fields{value: time.Time{}}, want: ""},
		{name: "return DateString", fields: fields{value: mockTime}, want: "2020-05-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Datetime{
				value: tt.fields.value,
			}
			if got := d.ToDatetimeString(); got != tt.want {
				t.Errorf("ToDatetimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
