package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestDMA(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s0 := pandas.NewSeriesWithoutType("f0", f0)
	fmt.Println(DMA(s0, 5))
	s2 := []float64{1, 2, 3, 4, 3, 3, 2, 1, stat.DTypeNaN, stat.DTypeNaN, stat.DTypeNaN, stat.DTypeNaN}
	fmt.Println(s2)
	//stat.Fill(s2, 1.0, true)
	//fmt.Println(s2)
	fmt.Println(DMA(s0, s2))
	csv := "~/.quant1x/data/cn/002528.csv"
	df := pandas.ReadCSV(csv)
	df.SetNames("data", "open", "close", "high", "low", "volume", "amount", "zf", "zdf", "zde", "hsl")
	fmt.Println(df)
	CLOSE := df.ColAsNDArray("close")

	cs := CLOSE
	REF10 := REF(CLOSE, 10)
	d1 := cs.Div(REF10)
	df01 := pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "x", d1)
	x0 := make([]stat.DType, CLOSE.Len())
	df01.Apply(func(idx int, v any) {
		f := v.(float32)
		t := stat.DType(0)
		if f >= 1.05 {
			t = stat.DType(1)
		}
		x0[idx] = t
	})
	n := BARSLAST2(pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "", x0))
	fmt.Println(n[len(n)-10:])
	x := DMA(CLOSE, pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "", n))

	//x := EMA(CLOSE, 7)
	sx := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "x", x)
	df = pandas.NewDataFrame(CLOSE, sx)
	fmt.Println(df)
}
