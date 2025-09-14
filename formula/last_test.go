package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/pandas"
)

func TestLAST(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12}
	i0 := CompareGt(f0, 10)
	s0 := pandas.NewSeriesWithoutType("f0", i0)
	fmt.Println(LAST(s0, 10, 5))
}
