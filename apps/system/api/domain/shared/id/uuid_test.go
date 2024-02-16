package id

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"reflect"
	"testing"
)

func TestUUID_ToFriendlyString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    UUID
		wantErr validation.Errors
	}{
		{name: "success", value: "aggvtvtpajybniivcqktozacgi", want: UUID{uuid.MustParse("018d59d6-6f02-7016-a115-141537640232")}, wantErr: nil},
		{name: "success", value: "aggvtv7lub6jhn6ahpxwmzl3vq", want: UUID{uuid.MustParse("018d59d7-eba0-7c93-b7c0-3bef66657bac")}, wantErr: nil},
		{name: "Empty", value: "", want: UUID{}, wantErr: test.NewValidationErrors(InvalidUUID, "").Errs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromFriendlyString(tt.value)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewFromFriendlyString() got = %v, want %v", got, tt.want)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewFromFriendlyString() got = %v, want %v", got, tt.want)
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewFromFriendlyString() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestNewUUIDFromString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    UUID
		wantErr validation.Errors
	}{
		{"success", "018d59d6-6f02-7016-a115-141537640232", UUID{uuid.MustParse("018d59d6-6f02-7016-a115-141537640232")}, nil},
		{"success", "018d9b4d-9438-79ac-b533-1323d4ec9b9f", UUID{uuid.MustParse("018d9b4d-9438-79ac-b533-1323d4ec9b9f")}, nil},
		{"empty", "", UUID{}, test.NewValidationErrors(InvalidUUID, "").Errs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUUIDFromString(tt.value)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewUUIDFromString() got = %v, want %v", got, tt.want)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewUUIDFromString() got = %v, want %v", got, tt.want)
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewUUIDFromString() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
