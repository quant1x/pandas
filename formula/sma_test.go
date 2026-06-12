package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/pandas"
)

func TestSMA(t *testing.T) {
	f0 := []float64{1, 2, 3, 4}
	fmt.Println(f0)
	s0 := pandas.NewSeriesWithoutType("f0", f0)
	fmt.Println(SMA(s0, 4, 1))
}
