package pandas

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSeriesString(t *testing.T) {
	s := NewSeriesWithType(reflect.String, "a", 1)
	fmt.Println(s)
}

func TestNewSeries(t *testing.T) {
	type args struct {
		typ    Type
		name   string
		values any
	}
	tests := []struct {
		name string
		args args
		want Series
	}{
		{
			name: "string",
			args: args{
				name:   "a",
				typ:    SERIES_TYPE_STRING,
				values: []float64{1},
			},
			want: NewSeriesWithType(reflect.String, "a", 1),
		},
		{
			name: "int64",
			args: args{
				name:   "a",
				typ:    SERIES_TYPE_INT64,
				values: []float64{1},
			},
			want: NewSeriesWithType(reflect.Int64, "a", 1),
		},
		{
			name: "float64",
			args: args{
				name:   "a",
				typ:    SERIES_TYPE_FLOAT64,
				values: []float64{1},
			},
			want: NewSeriesWithType(reflect.Float64, "a", 1),
		},
		{
			name: "float64<-string",
			args: args{
				name:   "x",
				typ:    SERIES_TYPE_FLOAT64,
				values: []string{"1"},
			},
			want: NewSeriesWithType(reflect.Float64, "x", 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeriesWithType(tt.args.typ, tt.args.name, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeriesWithType() = %v(%T), want %v(%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
