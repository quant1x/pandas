package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// MIN 两个序列横向对比
func MIN(S1, S2 stat.Series) stat.Series {
	d := stat.Minimum(S1.Floats(), S2.Floats())
	return pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "", d)

}
