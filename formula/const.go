package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// CONST 取S最后的值为常量
func CONST(S stat.Series) stat.Series {
	length := S.Len()
	s := S.Floats()
	s = stat.Repeat(s[length-1], S.Len())
	return pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "", s)
}
