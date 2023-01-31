package pandas

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSeriesFrame(t *testing.T) {
	type args struct {
		name string
		vals []interface{}
	}
	tests := []struct {
		name string
		args args
		want *SeriesFrame
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeriesFrame(tt.args.name, tt.args.vals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeriesFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeriesFrame(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := NewSeriesFrame("x", data)
	fmt.Printf("%+v\n", s1)

	var d1 any
	d1 = data
	s2 := NewSeriesFrame("x", d1)
	fmt.Printf("%+v\n", s2)
}
