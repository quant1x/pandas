package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/pandas"
)

func TestSQRT(t *testing.T) {
	f1 := []float32{1.1, 2.2, 1.3, 1.4}
	f2 := []float64{70, 80, 75, 83, 86}
	s1 := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "s1", f1)
	s2 := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "s2", f2)
	fmt.Println(SQRT(s1))
	fmt.Println(SQRT(s2))
}
