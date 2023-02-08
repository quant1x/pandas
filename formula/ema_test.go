package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"github.com/viterin/vek/vek32"
	"testing"
)

func TestEMA(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s0 := pandas.NewSeriesWithoutType("f0", f0)
	fmt.Println(EMA(s0, 7))
	csv := "~/.quant1x/data/cn/002528.csv"
	df := pandas.ReadCSV(csv)
	df.SetNames("data", "open", "close", "high", "low", "volume", "amount", "zf", "zdf", "zde", "hsl")
	fmt.Println(df)
	//df.SetName("收盘", "close")
	CLOSE := df.Col("close")

	cs := CLOSE.Values().([]float32)
	REF10 := REF(CLOSE, 10).([]float32)
	v1 := vek32.Div(cs, REF10)
	df01 := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "x", v1)
	x0 := make([]stat.DType, CLOSE.Len())
	df01.Apply(func(idx int, v any) {
		f := v.(float32)
		t := stat.DType(0)
		if f >= 1.05 {
			t = stat.DType(1)
		}
		x0[idx] = t
	})
	//x := stat.Where(v2, as, bs)
	n := BARSLAST(pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", x0))
	fmt.Println(n[len(n)-10:])
	x := EMA(CLOSE, pandas.NewSeries(pandas.SERIES_TYPE_DTYPE, "", n))

	//x := EMA(CLOSE, 7)
	sx := pandas.NewSeries(pandas.SERIES_TYPE_DTYPE, "x", x)
	df = pandas.NewDataFrame(CLOSE, sx)
	fmt.Println(df)

}
