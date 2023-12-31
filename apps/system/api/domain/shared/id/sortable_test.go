package id

import (
	"github.com/ryo034/react-go-template/packages/go/domain/shared/validation"
	"github.com/ryo034/react-go-template/packages/go/util/test"
	"reflect"
	"testing"
)

func TestNewSortableID_OK(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    SortableID
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{v: "01D3XQYXQZQZQZQZQZQZQZQZQZ"},
			want:    SortableID{v: "01D3XQYXQZQZQZQZQZQZQZQZQZ"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSortableID(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSortableID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSortableID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSortableID_Validate(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    SortableID
		wantErr validation.Errors
	}{
		{"required", args{v: ""}, SortableID{v: ""}, test.NewValidationErrors(InvalidSortableID, "").Errs},
		{"regex only Number", args{v: "0901234123400"}, SortableID{v: "0901234123400"}, test.NewValidationErrors(InvalidSortableID, "0901234123400").Errs},
		{"regex invalid ID", args{v: "ABCDEFG-12345"}, SortableID{v: "ABCDEFG-12345"}, test.NewValidationErrors(InvalidSortableID, "ABCDEFG-12345").Errs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSortableID(tt.args.v)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewSortableID() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewSortableID() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
