package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestBARSLAST(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12}
	i0 := CompareGt(f0, 100)
	s0 := stat.NewSeries[bool](i0...)
	fmt.Println(BARSLAST(s0))
}
