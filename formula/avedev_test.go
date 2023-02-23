package formula

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestAVEDEV(t *testing.T) {
	y := []float64{1.0, 0.41, 0.50, 0.61, 0.91, 2.02, 2.46}
	Y := pandas.NewSeriesWithoutType("y", y)
	fmt.Println(AVEDEV(Y, 5))

	df := stock.KLine("sz002528")
	df.SetNames("data", "open", "close", "high", "low", "volume", "amount", "zf", "zdf", "zde", "hsl")
	fmt.Println(df)
	fmt.Println(df.Names())
	df.SetName("收盘", "close")
	CLOSE := df.ColAsNDArray("close")

	c1 := AVEDEV(CLOSE, 5)
	C := pandas.NewSeriesWithoutType("c1", c1)
	df = pandas.NewDataFrame(C)
	fmt.Println(df)
}
