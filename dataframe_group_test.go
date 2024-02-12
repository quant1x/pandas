package pandas

import (
	"fmt"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestDataFrame_Group(t *testing.T) {
	type testStruct struct {
		A string
		B int
		C bool
		D float64
	}
	data := []testStruct{
		{"a", 1, true, 0.0},
		{"b", 2, false, 0.5},
		{"b", 3, false, 3.5},
		{"b", 4, false, 2.5},
		{"b", 5, false, 1.5},
		{"a", 6, true, 0.0},
		{"a", 7, true, 0.0},
		{"a", 8, true, 0.0},
		{"a", 9, true, 0.0},
		{"a", 10, true, 0.0},
		{"a", 11, true, 0.0},
	}
	df := LoadStructs(data)
	fmt.Println(df)

	df1 := df.Group("A", func(kind stat.Type, e any) bool {
		v := num.AnyToString(e)
		if v == "b" {
			return true
		}
		return false
	})
	fmt.Println(df1)
}
