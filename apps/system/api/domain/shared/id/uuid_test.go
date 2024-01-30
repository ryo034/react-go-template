package id

import (
	"testing"
)

func TestUUID_ToFriendlyString_OK(t *testing.T) {
	type fields struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "success",
			fields: fields{v: "018d59d6-6f02-7016-a115-141537640232"},
			want:   "aggvtvtpajybniivcqktozacgi",
		},
		{
			name:   "success",
			fields: fields{v: "018d59d7-eba0-7c93-b7c0-3bef66657bac"},
			want:   "aggvtv7lub6jhn6ahpxwmzl3vq",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := NewUUIDFromString(tt.fields.v)
			if got := u.ToFriendlyString(); got != tt.want {
				t.Errorf("ToFriendlyString() = %v, want %v", got, tt.want)
			}

			fs, _ := NewFromFriendlyString(tt.want)
			if got := fs.ToString(); got != tt.fields.v {
				t.Errorf("String() = %v, want %v", got, tt.fields.v)
			}
		})
	}
}
