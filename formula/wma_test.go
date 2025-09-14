package formula

import (
	"fmt"
	"testing"

	"gitee.com/quant1x/pandas"
)

func TestWMA(t *testing.T) {
	csv := "~/.quant1x/data/cn/002528.csv"
	df := pandas.ReadCSV(csv)
	fmt.Println(df)
	var (
		CLOSE = df.ColAsNDArray("close")
		HIGH  = df.ColAsNDArray("high")
		LOW   = df.ColAsNDArray("low")
		//VOL   = df.ColAsNDArray("volume")
		//DATE  = df.ColAsNDArray("date")
	)
	//length := CLOSE.Len()
	//N1 := N
	N2 := 5
	N3 := 2
	//N1:=6;
	//重心:(2*C+H+L)/4,COLOR00FFFF,LINETHICK0;
	ZX := CLOSE.Mul(2).Add(HIGH).Add(LOW).Div(4)
	//SJ:=WMA((重心-LLV(L,5))/(HHV(H,5)-LLV(L,5))*100,2);
	LLV5 := LLV(LOW, N2)
	HHV5 := HHV(HIGH, N2)

	sj1 := ZX.Sub(LLV5)
	sj2 := HHV5.Sub(LLV5)
	sj3 := sj1.Div(sj2).Mul(100)
	SJ := WMA(sj3, N3)
	//CLOSE := df.Col("close")
	fmt.Println(SJ)

}
