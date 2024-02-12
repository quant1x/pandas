package formula

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestBARSLASTS(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12}
	f0 = []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12, 0}
	f0 = []float64{4, 5, 6, 0, 8, 9, 10, 11, 12, 0}
	fmt.Println(f0)
	i0 := CompareGt(f0, 3)
	s0 := stat.NewSeries[bool](i0...)
	v := BARSLASTS(s0, 3)
	fmt.Println(v)
	df := pandas.NewDataFrame(stat.NewSeries[num.DType](f0...), v)
	fmt.Println(df)
}
