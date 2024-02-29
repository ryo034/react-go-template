package multi_factor

import (
	"reflect"
	"testing"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
)

func Test_multiFactors_Latest_OK(t *testing.T) {
	ph1, _ := phone.NewInternationalPhoneNumber("09000000000", "")
	dt1, _ := datetime.NewDatetimeFromString("2020-01-01 00:00:00")
	ph2, _ := phone.NewInternationalPhoneNumber("09000000001", "")
	dt2, _ := datetime.NewDatetimeFromString("2020-01-02 00:00:00")
	expected := NewMultiFactor(ph2, dt2)
	type fields struct {
		wrapped []MultiFactor
	}
	tests := []struct {
		name   string
		fields fields
		want   *MultiFactor
	}{
		{"success 1", fields{wrapped: []MultiFactor{NewMultiFactor(ph1, dt1), NewMultiFactor(ph2, dt2)}}, &expected},
		{"success 2", fields{wrapped: []MultiFactor{NewMultiFactor(ph2, dt2), NewMultiFactor(ph1, dt1)}}, &expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mfs := &multiFactors{
				wrapped: tt.fields.wrapped,
			}
			if got := mfs.Latest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Latest() = %v, want %v", got, tt.want)
			}
		})
	}
}
