package v1

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestDataFrame_Remove(t *testing.T) {
	type testStruct struct {
		A string
		B int
		C bool
		D float64
	}
	data := []testStruct{
		{"a", 1, true, 0.0},
		{"b", 2, false, 0.5},
	}
	df1 := LoadStructs(data)
	fmt.Println(df1)

	// 增加1列
	s_e := GenericSeries[string]("x", "a0", "a1", "a2", "a3", "a4")
	df2 := df1.Join(s_e)
	fmt.Println(df2)
	r := stat.RangeFinite(3, 3)
	df3 := df2.Remove(r)
	fmt.Println(df3)

}
