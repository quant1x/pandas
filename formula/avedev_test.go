package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestAVEDEV(t *testing.T) {
	y := []float64{1.0, 0.41, 0.50, 0.61, 0.91, 2.02, 2.46}
	Y := pandas.NewSeriesWithoutType("y", y)
	fmt.Println(AVEDEV(Y, 5))
}
