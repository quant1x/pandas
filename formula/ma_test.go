package formula

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestMA(t *testing.T) {
	type testStruct struct {
		A string
		B int
		C bool
		D float32
	}
	data := []testStruct{
		{"a", 1, true, 0.0},
		{"b", 2, false, 0.5},
	}
	df1 := pandas.LoadStructs(data)
	fmt.Println(df1)
	// 修改列名
	_ = df1.SetNames("a", "b", "c", "d")
	// 增加1列
	s_e := pandas.GenericSeries[string]("", "a0", "a1", "a2", "a3")
	df2 := df1.Join(s_e)
	fmt.Println(df2)
	B := df2.Col("b")

	// 2日均线
	r2 := MAV1(B, 2)
	fmt.Println(r2)
}

func TestMAx(t *testing.T) {
	s := stat.NewSeries(1, 2, 3, 4, 5, 6)
	ma5 := MA(s, 5)
	fmt.Println(ma5)
	arr := ma5.DTypes()
	old := arr[len(arr)-2]
	new := s.IndexOf(-1)
	n := MAx(5, old, num.AnyToFloat64(new))
	fmt.Println(n)
}
