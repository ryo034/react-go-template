package account

import (
	"reflect"
	"testing"
)

func TestNewName(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Name
		wantErr bool
	}{
		{"normal", args{"test"}, Name{"test"}, false},
		{"", args{""}, Name{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
