package phone

import (
	"reflect"
	"testing"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
)

func TestNewInternationalPhoneNumber_Validate(t *testing.T) {
	type args struct {
		v      string
		region string
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr validation.Errors
	}{
		{name: "required", args: args{v: "", region: "JP"}, want: Number{value: ""}, wantErr: test.NewValidationErrors(InvalidPhoneNumber, "").Errs},
		{"regex", args{v: "0901234123400", region: "JP"}, Number{value: "0901234123400"}, test.NewValidationErrors(InvalidPhoneNumber, "0901234123400").Errs},
		{"regex", args{v: "109012341234", region: "JP"}, Number{value: "109012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "109012341234").Errs},
		{"regex", args{v: "00012341234", region: "JP"}, Number{value: "00012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "00012341234").Errs},
		{"regex", args{v: "01012341234", region: "JP"}, Number{value: "01012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "01012341234").Errs},
		{"regex", args{v: "02012341234", region: "JP"}, Number{value: "02012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "02012341234").Errs},
		{"regex", args{v: "03012341234", region: "JP"}, Number{value: "03012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "03012341234").Errs},
		{"regex", args{v: "04012341234", region: "JP"}, Number{value: "04012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "04012341234").Errs},
		{"regex", args{v: "05012341234", region: "JP"}, Number{value: "05012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "05012341234").Errs},
		{"regex", args{v: "06012341234", region: "JP"}, Number{value: "06012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "06012341234").Errs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInternationalPhoneNumber(tt.args.v, tt.args.region)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewInternationalPhoneNumber() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewInternationalPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestNewInternationalPhoneNumber_OK(t *testing.T) {
	type args struct {
		v      string
		region string
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr bool
	}{
		{"regex", args{v: "+819012341234", region: "JP"}, Number{value: "+819012341234", region: "JP"}, false},
		{"no region", args{v: "+819012341234", region: ""}, Number{value: "+819012341234", region: "JP"}, false},
		{"local number", args{v: "09012341234", region: "JP"}, Number{value: "+819012341234", region: "JP"}, false},
		{"no region local number", args{v: "09012341234", region: ""}, Number{value: "+819012341234", region: "JP"}, false},
		{"no region hyphen number", args{v: "090-1234-1234", region: ""}, Number{value: "+819012341234", region: "JP"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInternationalPhoneNumber(tt.args.v, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInternationalPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternationalPhoneNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}
