package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestABS(t *testing.T) {
	d1 := []int32{1, -1, 2, -2}
	s := pandas.NewSeries(stat.SERIES_TYPE_FLOAT64, "", d1)
	fmt.Println(ABS(s))

}
