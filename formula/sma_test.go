package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"github.com/viterin/vek/vek32"
	"testing"
)

func TestSMA(t *testing.T) {
	f0 := []float64{1, 2, 3, 4}
	s0 := pandas.NewSeriesWithoutType("f0", f0)
	fmt.Println(SMA(s0, 4, 1))
	csv := "~/.quant1x/data/cn/002528.csv"
	df := pandas.ReadCSV(csv)
	df.SetNames("data", "open", "close", "high", "low", "volume", "amount", "zf", "zdf", "zde", "hsl")
	fmt.Println(df)
	fmt.Println(df.Names())
	df.SetName("收盘", "close")
	CLOSE := df.Col("close")
	cs := CLOSE.Values().([]float32)
	REF10 := REF(CLOSE, 10).([]float32)
	v1 := vek32.Div(cs, REF10)
	//as := stat.Repeat[float32](1, CLOSE.Len())
	//bs := stat.Repeat[float32](0, CLOSE.Len())
	df01 := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "x", v1)
	x := make([]float32, CLOSE.Len())
	df01.Apply(func(idx int, v any) {
		f := v.(float32)
		t := float32(0)
		if f >= 1.05 {
			t = float32(1)
		}
		x[idx] = t
	})
	//x := stat.Where(v2, as, bs)
	n := barslast(pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", x))
	fmt.Println(n[len(n)-10:])
	//r1 := SMA(CLOSE, pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", n), 1)
	r1 := SMA(CLOSE, 6, 1)
	s2 := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "sma", r1)
	df2 := pandas.NewDataFrame(s2)
	fmt.Println(df2)
}
