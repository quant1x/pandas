package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestFORCAST(t *testing.T) {
	x := []float64{1.0, 0.41, 0.50, 0.61, 0.91, 2.02, 2.46}
	X := pandas.NewSeriesWithoutType("x", x)
	fmt.Println(FORCAST(X, 5))
}
