package formula

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/num/labs"
	"gitee.com/quant1x/pandas"
	"slices"
	"testing"
)

func TestSUM_basic(t *testing.T) {
	f0 := []float64{1.1, 2.2, 1.3, 1.4}
	f1 := []float64{70, 80, 75, 83, 86}
	f2 := []float64{90, 69, 60, 88, 87}

	s0 := pandas.NewSeriesWithoutType("f0", f0)
	s1 := pandas.NewSeriesWithoutType("f1", f1)
	s2 := pandas.NewSeriesWithoutType("f2", f2)
	fmt.Println(SUM(s0, 4))
	fmt.Println(v1SUM(s1, 5))
	fmt.Println(v2SUM(s2, 5))
}

func TestSUM_release(t *testing.T) {
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
			name: "float64",
			args: args{
				S: pandas.ToVector[float64](1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				N: 5,
			},
			want: pandas.ToVector[float64](num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), 15, 20, 25, 30, 35, 40),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SUM(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("SUM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSUM_v1(t *testing.T) {
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
			name: "float64",
			args: args{
				S: pandas.ToVector[float64](1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				N: 5,
			},
			want: pandas.ToVector[float64](num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), 15, 20, 25, 30, 35, 40),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := v1SUM(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("v1SUM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSUM_v2(t *testing.T) {
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
			name: "float64",
			args: args{
				S: pandas.ToVector[float64](1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				N: 5,
			},
			want: pandas.ToVector[float64](num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), 15, 20, 25, 30, 35, 40),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := v2SUM(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("v2SUM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSUM_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkSUM_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		SUM(s, 10)
	}
}

func BenchmarkSUM_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v1SUM(s, 10)
	}
}

func BenchmarkSUM_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v2SUM(s, 10)
	}
}
