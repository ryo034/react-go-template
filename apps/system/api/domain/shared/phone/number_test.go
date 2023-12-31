package phone

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"reflect"
	"testing"
)

func TestNewPhoneNumber_OK(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr bool
	}{
		{"success", args{v: "07012341234"}, Number{value: "07012341234"}, false},
		{"success", args{v: "08012341234"}, Number{value: "08012341234"}, false},
		{"success", args{v: "09012341234"}, Number{value: "09012341234"}, false},
		{"success", args{v: "09000000000"}, Number{value: "09000000000"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPhoneNumber(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhoneNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPhoneNumber_Validate(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr validation.Errors
	}{
		{name: "required", args: args{v: ""}, want: Number{value: ""}, wantErr: test.NewValidationErrors(InvalidPhoneNumber, "").Errs},
		{"regex", args{v: "0901234123400"}, Number{value: "0901234123400"}, test.NewValidationErrors(InvalidPhoneNumber, "0901234123400").Errs},
		{"regex", args{v: "109012341234"}, Number{value: "109012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "109012341234").Errs},
		{"regex", args{v: "00012341234"}, Number{value: "00012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "00012341234").Errs},
		{"regex", args{v: "01012341234"}, Number{value: "01012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "01012341234").Errs},
		{"regex", args{v: "02012341234"}, Number{value: "02012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "02012341234").Errs},
		{"regex", args{v: "03012341234"}, Number{value: "03012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "03012341234").Errs},
		{"regex", args{v: "04012341234"}, Number{value: "04012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "04012341234").Errs},
		{"regex", args{v: "05012341234"}, Number{value: "05012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "05012341234").Errs},
		{"regex", args{v: "06012341234"}, Number{value: "06012341234"}, test.NewValidationErrors(InvalidPhoneNumber, "06012341234").Errs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPhoneNumber(tt.args.v)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewPhoneNumber() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestPhoneNumber_ToInternationalNumberString_OK(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"success", fields{value: "07012341234"}, "+817012341234"},
		{"success", fields{value: "08012341234"}, "+818012341234"},
		{"success", fields{value: "09012341234"}, "+819012341234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Number{
				value: tt.fields.value,
			}
			if got := e.ToInternationalNumberString(); got != tt.want {
				t.Errorf("ToInternationalNumberString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternationalPhoneNumber(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr bool
	}{
		{"required", args{v: ""}, Number{value: ""}, true},
		{"regex", args{v: "+819012341234"}, Number{value: "09012341234"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInternationalPhoneNumber(tt.args.v)
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
