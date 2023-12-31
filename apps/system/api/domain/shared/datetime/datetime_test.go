package datetime

import (
	"reflect"
	"testing"
	"time"
)

var now = time.Now().UTC()

func Test_Datetime_ToTime_OK(t *testing.T) {
	type fields struct {
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "Timeで返す",
			fields: fields{
				value: now,
			},
			want: now,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Datetime{
				value: tt.fields.value,
			}
			if got := d.ToTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Datetime_ToDateString_OK(t *testing.T) {
	mockTime := time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC)

	type fields struct {
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "空の場合に空文字が返ってくる",
			fields: fields{
				value: time.Time{},
			},
			want: "",
		},
		{
			name: "文字列が返ってくる",
			fields: fields{
				value: mockTime,
			},
			want: "2020-05-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Datetime{
				value: tt.fields.value,
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
		{
			name: "空の場合に空文字が返ってくる",
			fields: fields{
				value: time.Time{},
			},
			want: "",
		},
		{
			name: "文字列が返ってくる",
			fields: fields{
				value: mockTime,
			},
			want: "2020-05-01 00:00:00",
		},
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
