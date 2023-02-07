package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// CONST 取S最后的值为常量
func CONST(S pandas.Series) pandas.Series {
	length := S.Len()
	s := S.Float()
	s = stat.Repeat(s[length-1], S.Len())
	return pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", s)
}
