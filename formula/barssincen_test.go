package formula

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestBARSSINCEN(t *testing.T) {
	f1 := []int64{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	s1 := pandas.NewSeriesWithType(pandas.SERIES_TYPE_DTYPE, "", f1)
	df := pandas.NewDataFrame(s1)
	fmt.Println(df)

	b1 := s1.Logic(func(idx int, v any) bool {
		f := v.(num.DType)
		return f > 3
	})
	df = df.Join(pandas.NewSeriesWithType(pandas.SERIES_TYPE_BOOL, "r", b1))
	fmt.Println(df)
	//c1 = df > 3
	r1 := BARSSINCEN(df.Col("r"), 4)
	fmt.Println(r1)
}
