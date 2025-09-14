package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/pandas"
)

func TestABS(t *testing.T) {
	d1 := []int32{1, -1, 2, -2}
	s := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "", d1)
	fmt.Println(s)
	s1 := ABS(s)
	fmt.Println(s1)
	fmt.Println(s)
}
