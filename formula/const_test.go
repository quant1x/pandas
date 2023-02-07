package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestCONST(t *testing.T) {
	f0 := []float64{1, 2, 3, 4}

	s0 := pandas.NewSeriesWithoutType("f0", f0)
	fmt.Println(CONST(s0))
}
