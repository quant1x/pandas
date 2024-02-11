package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// ABS 计算S的绝对值
func ABS(S stat.Series) stat.Series {
	s := S.DTypes()
	d := stat.Abs(s)
	//fmt.Printf("%p\n", d)
	//return pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "", d)
	//return stat.NewSeries(d...)
	return stat.ToSeries(d...)
}
