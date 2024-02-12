package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestMIN(t *testing.T) {
	f1 := []float32{1.1, 2.2, 1.3, 1.4}
	f2 := []float32{1.2, 1.2, 3.3}
	s1 := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT64, "x1", f1)
	s2 := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT64, "x2", f2)
	fmt.Println(MIN(s1, s2))
}
