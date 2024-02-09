package account

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/id"
	"reflect"
	"testing"
)

func TestNewID(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    ID
		wantErr bool
	}{
		{"normal", args{"018d8f91-dc87-7b94-914d-6a5b03dc4a9f"}, ID{id.NewUUID(uuid.MustParse("018d8f91-dc87-7b94-914d-6a5b03dc4a9f"))}, false},
		{"invalid", args{"invalid"}, ID{}, true},
		{"", args{""}, ID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewID(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
