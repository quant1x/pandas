package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// MIN 两个序列横向对比
func MIN(S1, S2 pandas.Series) pandas.Series {
	d := stat.Minimum(S1.Float(), S2.Float())
	return pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", d)

}
