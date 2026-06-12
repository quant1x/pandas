package formula

import (
	"fmt"
	"slices"
	"testing"

	"github.com/quant1x/num/labs"
	"github.com/quant1x/pandas"
)

func TestEmaIncr_basic(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := pandas.NewSeries[float64](f0...)
	v0 := EMA(s, 7)
	fmt.Println(v0)
	v1 := EMA(s, 7)
	fmt.Println(v1)
	last := v1.IndexOf(-2).(float64)
	alpha := float64(2) / float64(1+7)
	//(1−α)*y(t−1) + α*x(t)
	//last = (beta * last) + (alpha * x)
	v2 := (1-alpha)*last + alpha*9
	fmt.Println(v2)
}

func TestEMA(t *testing.T) {
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
			name: "ema-1",
			args: args{
				S: pandas.NewSeries[float64](1, 2, 3, 4),
				N: 4,
			},
			want: pandas.NewSeries[float64](1., 1.4, 2.04, 2.824),
		},
		{
			name: "ema-2",
			args: args{
				S: pandas.NewSeries[float64](1, 2, 3, 4, 5, 6, 7, 8, 9),
				N: 7,
			},
			want: pandas.NewSeries[float64](1, 1.25, 1.6875, 2.265625, 2.94921875, 3.7119140625, 4.533935546875, 5.40045166015625, 6.3003387451171875),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EMA(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("EMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkEMA_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkEMA_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		EMA(s, 10)
	}
}

func BenchmarkEMA_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v1EMA(s, 10)
	}
}

func BenchmarkEMA_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v2EMA(s, 10)
	}
}
