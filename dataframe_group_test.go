package pandas

import (
	"fmt"
	"testing"

	"github.com/quant1x/num"
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

	df1 := df.Group("A", func(kind Type, e any) bool {
		v := num.AnyToString(e)
		if v == "b" {
			return true
		}
		return false
	})
	fmt.Println(df1)
}
