package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/pandas"
)

func TestSLOPE(t *testing.T) {
	//x := []float64{0.0, 0.1, 0.2, 0.3, 0.5, 0.8, 1.0}
	y := []float64{1.0, 0.41, 0.50, 0.61, 0.91, 2.02, 2.46}
	//X := pandas.NewSeriesWithoutType("x", x)
	Y := pandas.NewSeriesWithoutType("y", y)
	fmt.Println(SLOPE(Y, 5))
}
