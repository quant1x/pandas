package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestABS(t *testing.T) {
	v1 := []int32{1, -1, 2, -2}
	s := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT64, "", v1)
	fmt.Println(ABS(s))

}
