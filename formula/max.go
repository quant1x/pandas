package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// MAX 两个序列横向对比
func MAX(S1, S2 stat.Series) stat.Series {
	d := stat.Maximum(S1.Floats(), S2.Floats())
	return pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "", d)

}
