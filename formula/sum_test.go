package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestSUM(t *testing.T) {
	f0 := []float64{1.1, 2.2, 1.3, 1.4}
	f1 := []float64{70, 80, 75, 83, 86}
	f2 := []float64{90, 69, 60, 88, 87}

	s0 := pandas.NewSeriesWithoutType("f0", f0)
	s1 := pandas.NewSeriesWithoutType("f1", f1)
	s2 := pandas.NewSeriesWithoutType("f2", f2)
	fmt.Println(SUM(s0, 4))
	fmt.Println(SUM(s1, 5))
	fmt.Println(SUM(s2, 5))
}
