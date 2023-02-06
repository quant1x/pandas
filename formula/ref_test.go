package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestREF(t *testing.T) {
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
	A := df2.Col("a")
	B := df2.Col("b")
	C := df2.Col("c")
	D := df2.Col("d")

	// 2日前的D值
	r2 := REF(D, 2)
	fmt.Println(r2)

	_ = A
	_ = B
	_ = C
	_ = D
}
