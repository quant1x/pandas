package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestBARSLASTCOUNT(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12}
	i0 := CompareGte(f0, 1)
	s0 := pandas.NewSeriesWithoutType("f0", i0)
	fmt.Println(BARSLASTCOUNT(s0))
}
