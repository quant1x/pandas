package pandas

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSeriesFrame(t *testing.T) {

	//sf := NewSeries(SERIES_TYPE_STRING, "x", []string{"1", "2", "3"})
	//fmt.Println(sf)

	type args struct {
		t    string
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
			if got := NewSeries(tt.args.t, tt.args.name, tt.args.vals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeriesFrame(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := NewSeries(SERIES_TYPE_FLOAT, "x", data)
	fmt.Printf("%+v\n", s1)

	var d1 any
	d1 = data
	s2 := NewSeries(SERIES_TYPE_FLOAT, "x", d1)
	fmt.Printf("%+v\n", s2)
}
