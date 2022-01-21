package sap_api_time_value_converter

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertToTime(t *testing.T) {
	type args struct {
		sapTime string
	}
	type testStr struct {
		name string
		args args
		want time.Time
	}
	tests := []testStr{
		func() testStr {
			now := time.Now()
			return testStr{
				name: "OK now time",
				args: args{
					sapTime: fmt.Sprintf(`\/Date(%d)\/`, now.UnixMilli()),
				},
				want: now,
			}
		}(),
		func() testStr {
			return testStr{
				name: "OK now time",
				args: args{
					sapTime: `\/Date(1642757478000)\/`,
				},
				want: time.Date(2022, 1, 21, 9, 31, 18, 0, time.UTC),
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToTimeFormat(tt.args.sapTime); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want.UnixMilli(), got.UnixMilli(), "not same time")
			}
		})
	}
}
