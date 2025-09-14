package formula

import (
	"slices"
	"testing"

	"github.com/quant1x/num"
	"github.com/quant1x/num/labs"
	"github.com/quant1x/pandas"
)

func TestHHV_basic(t *testing.T) {

}

func TestHHV(t *testing.T) {
	type args struct {
		S pandas.Series
		N any
	}
	tests := []struct {
		name string
		args args
		want pandas.Series
	}{
		{
			name: "hhv",
			args: args{
				S: pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "x", []float64{2, 1, 3, 4, 5, 2}),
				N: 3,
			},
			want: pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "x", []float64{num.Float64NaN(), num.Float64NaN(), 3, 4, 5, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HHV(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("HHV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHHV_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkHHV_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		HHV(s, 10)
	}
}

func BenchmarkHHV_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v1HHV(s, 10)
	}
}

func BenchmarkHHV_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v2HHV(s, 10)
	}
}
