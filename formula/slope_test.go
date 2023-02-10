package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestSLOPE(t *testing.T) {
	x := []float64{0.0, 0.1, 0.2, 0.3, 0.5, 0.8, 1.0}
	y := []float64{1.0, 0.41, 0.50, 0.61, 0.91, 2.02, 2.46}
	X := pandas.NewSeriesWithoutType("x", x)
	Y := pandas.NewSeriesWithoutType("y", y)
	fmt.Println(SLOPE(Y, 5))

	csv := "~/.quant1x/data/cn/002528.csv"
	df := pandas.ReadCSV(csv)
	df.SetNames("data", "open", "close", "high", "low", "volume", "amount", "zf", "zdf", "zde", "hsl")
	fmt.Println(df)
	fmt.Println(df.Names())
	df.SetName("收盘", "close")
	CLOSE := df.Col("close")

	c1 := SLOPE(CLOSE, 5)
	C := pandas.NewSeriesWithoutType("c1", c1)
	df = pandas.NewDataFrame(C)
	fmt.Println(df)

	_ = X
}
