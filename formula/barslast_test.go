package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestBARSLAST(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12}
	i0 := CompareGt(f0, 10)
	s0 := pandas.SliceToSeries(i0)
	fmt.Println(BARSLAST(s0))
}
