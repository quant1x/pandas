package internal

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/pandas"
)

// KLine 加载K线
func KLine(code string) pandas.DataFrame {
	filename := cache.GetKLineFilename(code)
	df := pandas.ReadCSV(filename)
	_ = df.SetNames("date", "open", "high", "low", "close", "volume")
	return df
}
