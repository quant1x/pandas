package pandas

import (
	"fmt"
	"testing"
)

func TestLoadStructs(t *testing.T) {
	type testStruct struct {
		A string
		B int
		C bool
		D float64
	}
	type testStructTags struct {
		A string  `dataframe:"a,string"`
		B int     `dataframe:"b,string"`
		C bool    `dataframe:"c,string"`
		D float64 `dataframe:"d,string"`
		E int     `dataframe:"-"` // ignored
		f int     // ignored
	}
	data := []testStruct{
		{"a", 1, true, 0.0},
		{"b", 2, false, 0.5},
	}
	dataTags := []testStructTags{
		{"a", 1, true, 0.0, 0, 0},
		{"NA", 2, false, 0.5, 1, 3},
		{"NA", 3, false, 1.5, 2, 4},
	}
	df1 := LoadStructs(data)
	fmt.Println(df1)
	df2 := LoadStructs(dataTags)
	fmt.Println(df2)
}
