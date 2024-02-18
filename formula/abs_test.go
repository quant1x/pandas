package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestABS(t *testing.T) {
	d1 := []int32{1, -1, 2, -2}
	s := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "", d1)
	fmt.Println(s)
	s1 := ABS(s)
	//fmt.Printf("%p\n", s1)
	fmt.Println(s1)
	fmt.Println(s)
}
