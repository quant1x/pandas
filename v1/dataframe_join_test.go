package v1

import (
	"fmt"
	"testing"
)

func TestDataFrame_Join(t *testing.T) {
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
	df1 := LoadStructs(data)
	fmt.Println(df1)

	// 增加1列
	s_e := GenericSeries[string]("", "a0", "a1", "a2", "a3")
	df2 := df1.Join(s_e)
	fmt.Println(df2)
}
