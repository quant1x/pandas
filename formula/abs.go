package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// ABS 计算S的绝对值
func ABS(S pandas.Series) pandas.Series {
	s := S.DTypes()
	d := num.Abs(s)
	return pandas.SliceToSeries(d)
}
