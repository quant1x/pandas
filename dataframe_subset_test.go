package pandas

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDataFrame_IndexOf(t *testing.T) {
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
	//fmt.Println(df1)

	// 增加1列
	s_e := GenericSeries[string]("x", "a0", "a1", "a2", "a3", "a4")
	df2 := df1.Join(s_e)
	//fmt.Println(df2)
	df := df2.Select([]string{"A"})
	fmt.Println(df)
	m := df.IndexOf(1, true)
	a, ok := m["A"]
	if ok {
		fmt.Println(a)
		if v, ok := a.(reflect.Value); ok {
			v.SetString("1")
		}
		fmt.Println(df)
	}
}
