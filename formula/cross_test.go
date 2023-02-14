package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestCROSS(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 8, 10}
	d2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := pandas.NewSeriesWithoutType("d1", d1)
	s2 := pandas.NewSeriesWithoutType("d2", d2)

	fmt.Println(CROSS(d1, d2))
	fmt.Println(CROSS1(s1, s2))
}
