package formula

import (
	"fmt"
	"gitee.com/quant1x/engine/datasource/base"
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
	s := stat.NewSeries[float64](f0...)
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

func TestEmaIncrV2(t *testing.T) {
	code := "002528"
	date := "20240126"
	klines := base.CheckoutKLines(code, date)
	if len(klines) == 0 {
		panic("no data")
	}
	df := pandas.LoadStructs(klines)
	if df.Nrow() == 0 {
		panic("加载k线失败")
	}
	var (
		DATE  = df.Col("date")
		CLOSE = df.ColAsNDArray("close")
	)

	//DIF:EMA(CLOSE,SHORT)-EMA(CLOSE,LONG);
	//DEA:EMA(DIF,MID)
	//MACD:(DIF-DEA)*2,COLORSTICK;
	SHORT := 12
	LONG := 26
	MID := 9

	d1 := EMA(CLOSE, SHORT)
	d2 := EMA(CLOSE, LONG)
	DIF := d1.Sub(d2)
	DEA := EMA(DIF, MID)
	MACD := DIF.Sub(DEA).Mul(2)

	df = pandas.NewDataFrame(
		DATE,
		pandas.NewSeriesWithoutType("d1", d1),
		pandas.NewSeriesWithoutType("d2", d2),
		pandas.NewSeriesWithoutType("DIF", DIF),
		pandas.NewSeriesWithoutType("DEA", DEA),
		pandas.NewSeriesWithoutType("MACD", MACD),
	)
	fmt.Println(df)
	fmt.Println("==============================================================================================================")
	lastClose := CLOSE.IndexOf(-1).(float64)
	fmt.Println("确定最新收盘价:", lastClose)
	df1 := df.IndexOf(-2)
	fmt.Println(df1)
	//date0 := df1["date"]
	dif1 := df1["d1"].(float64)
	dif2 := df1["d2"].(float64)
	dif1 = EmaIncr(lastClose, dif1, AlphaOfEMA(SHORT))
	dif2 = EmaIncr(lastClose, dif2, AlphaOfEMA(LONG))
	dif := dif1 - dif2
	lastDif := df1["DIF"].(float64)
	lastDea := df1["DEA"].(float64)
	fmt.Println("lastDif", lastDif)
	alpha := AlphaOfEMA(MID)
	fmt.Println("dea-alpha:", alpha)
	//t1 := -0.446
	//fmt.Println("xx:", (1-alpha)*lastDif+alpha*dif)
	dea := EmaIncr(dif, lastDea, AlphaOfEMA(MID))
	macd := (dif - dea) * 2

	fmt.Println("date:", DATE.IndexOf(-1))
	fmt.Println(" dif:", dif)
	fmt.Println(" dea:", dea)
	fmt.Println("macd:", macd)
}
