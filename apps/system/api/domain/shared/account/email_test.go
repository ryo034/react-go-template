package account

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"reflect"
	"testing"
)

func TestNewEmail_OK(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Email
		wantErr bool
	}{
		{
			name:    "OK",
			args:    args{v: "test@example.com"},
			want:    Email{value: "test@example.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmail(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmail_Validate(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Email
		wantErr validation.Errors
	}{
		{
			name:    "required",
			args:    args{v: ""},
			want:    Email{value: ""},
			wantErr: test.NewValidationErrors(InvalidEmail, "").Errs,
		},
		{
			name:    "required @",
			args:    args{v: "testexample.com"},
			want:    Email{value: "testexample.com"},
			wantErr: test.NewValidationErrors(InvalidEmail, "testexample.com").Errs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmail(tt.args.v)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewEmail() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
