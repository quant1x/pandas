package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

func MAX(S1, S2 pandas.Series) pandas.Series {
	d := stat.Maxinum(S1.Float(), S2.Float())
	return pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", d)

}
