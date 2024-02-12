package pandas

import (
	"fmt"
	"gitee.com/quant1x/num"
	"testing"
)

func TestNDFrame_Diff(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDFrame[float64]("x", d1...)
	df := NewDataFrame(s1)
	fmt.Println(df)
	fmt.Println("------------------------------------------------------------")
	N := 1
	fmt.Println("固定的参数, N =", N)
	r1 := df.Col("x").Diff(N).Values()
	fmt.Println("序列化结果:", r1)
	fmt.Println("------------------------------------------------------------")
	d2 := []float64{1, 2, 3, 4, 3, 3, 2, 1, num.Nil2Float64, num.Nil2Float64, num.Nil2Float64, num.Nil2Float64}
	s2 := NewSeries(SERIES_TYPE_FLOAT64, "x", d2)
	fmt.Printf("序列化参数: %+v\n", s2.Values())
	r2 := df.Col("x").Diff(s2).Values()
	fmt.Println("序列化结果:", r2)
}
