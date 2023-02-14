package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"math"
	"testing"
)

func TestMAX(t *testing.T) {
	fmt.Println(float64(1.4) > math.NaN())
	fmt.Println(float64(1.4) < math.NaN())
	f1 := []float32{1.1, 2.2, 1.3, 1.4}
	f2 := []float32{1.2, 1.2, 3.3}
	s1 := pandas.NewSeries(stat.SERIES_TYPE_FLOAT64, "x1", f1)
	s2 := pandas.NewSeries(stat.SERIES_TYPE_FLOAT64, "x2", f2)
	fmt.Println(MAX(s1, s2))
}
