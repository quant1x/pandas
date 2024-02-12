package pandas

import (
	"fmt"
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
	fmt.Println(df1)
}
