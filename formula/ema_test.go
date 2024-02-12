package formula

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
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
	CLOSE := df.ColAsNDArray("close")

	cs := CLOSE
	REF10 := REF(CLOSE, 10)
	d1 := cs.Div(REF10)
	df01 := pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "x", d1)
	x0 := make([]num.DType, CLOSE.Len())
	df01.Apply(func(idx int, v any) {
		f := v.(float32)
		t := num.DType(0)
		if f >= 1.05 {
			t = num.DType(1)
		}
		x0[idx] = t
	})
	//x := stat.Where(v2, as, bs)
	n := BARSLAST2(pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "", x0))
	fmt.Println(n[len(n)-10:])
	x := EMA(CLOSE, pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "", n))

	//x := EMA(CLOSE, 7)
	sx := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "x", x)
	df = pandas.NewDataFrame(CLOSE, sx)
	fmt.Println(df)
}

func TestEmaIncr(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := stat.NewNDArray[float64](f0...)
	v0 := EMA(s, 7)
	fmt.Println(v0)
	v1 := EMA(s, 7)
	fmt.Println(v1)
	last := v1.IndexOf(-2).(float64)
	alpha := float64(2) / float64(1+7)
	//(1−α)*y(t−1) + α*x(t)
	//last = (beta * last) + (alpha * x)
	v2 := (1-alpha)*last + alpha*9
	fmt.Println(v2)
}
