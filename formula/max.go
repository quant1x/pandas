package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// MAX 两个序列横向对比
func MAX(S1, S2 pandas.Series) pandas.Series {
	d := stat.Maximum(S1.Float(), S2.Float())
	return pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", d)

}
