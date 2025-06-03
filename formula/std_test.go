package formula

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/num/labs"
	"gitee.com/quant1x/pandas"
	"slices"
	"testing"
)

func TestSTD_basic(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	s := pandas.Vector(x)
	r1 := STD(s, 5)
	fmt.Println(r1)

	r2 := num.RollingV1(x, 5, func(N num.DType, values ...float64) float64 {
		fmt.Println(values)
		return num.Std(values)
	})
	fmt.Println(r2)

	fmt.Println(labs.DeepEqual(r1.Values(), r2))
}

func TestSTD_basic_dynamic(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	s := pandas.Vector(x)
	w := []int{5, 5, 5, 5, 5, 100, 5, 5, 5, 5}
	r1 := STD(s, w)
	fmt.Println(r1)

	r2 := num.RollingV1(x, 5, func(N num.DType, values ...float64) float64 {
		fmt.Println(values)
		return num.Std(values)
	})
	fmt.Println(r2)

	fmt.Println(labs.DeepEqual(r1.Values(), r2))
}

func TestSTD(t *testing.T) {
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
			want: pandas.ToVector[float64](num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), 1.4142135623730951, 1.4142135623730951, 1.4142135623730951, 1.4142135623730951, 1.4142135623730951, 1.4142135623730951),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := STD(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("STD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSTD_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkSTD_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		STD(s, 20)
	}
}

func BenchmarkSTD_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v1STD(s, 10)
	}
}

func BenchmarkSTD_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v2STD(s, 10)
	}
}
