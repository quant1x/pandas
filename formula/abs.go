package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

func ABS(S pandas.Series) pandas.Series {
	s := S.DTypes()
	d := stat.Abs(s)
	return pandas.NewSeries(pandas.SERIES_TYPE_DTYPE, "", d)
}
