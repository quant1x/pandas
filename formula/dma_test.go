package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

func TestDMA(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s0 := pandas.NewSeriesWithoutType("f0", f0)
	fmt.Println(DMA(s0, 5))
	s2 := []float64{1, 2, 3, 4, 3, 3, 2, 1, num.NaN(), num.NaN(), num.NaN(), num.NaN()}
	fmt.Println(s2)
}
