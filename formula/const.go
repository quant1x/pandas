package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// CONST 取S最后的值为常量
func CONST(S pandas.Series) pandas.Series {
	length := S.Len()
	s := S.Floats()
	s = num.Repeat(s[length-1], S.Len())
	return pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", s)
}
