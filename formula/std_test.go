package formula

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/num/labs"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestSTD_basic(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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

func TestSTD(t *testing.T) {
	f0 := []float64{1.1, 2.2, 1.3, 1.4}
	f1 := []float64{70, 80, 75, 83, 86}
	f2 := []float64{90, 69, 60, 88, 87}

	s0 := pandas.NewSeriesWithoutType("f0", f0)
	s1 := pandas.NewSeriesWithoutType("f1", f1)
	s2 := pandas.NewSeriesWithoutType("f2", f2)
	fmt.Println(STD(s0, 4))
	fmt.Println(STD(s1, 5))
	fmt.Println(STD(s2, 5))
}

func BenchmarkSTD_release(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s := pandas.Vector(x)
		r := STD(s, 5)
		_ = r
	}
}

func BenchmarkSTD_v3Rolling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		num.RollingV1(x, 5, func(N num.DType, values ...float64) float64 {
			return num.Std(values)
		})
	}
}
