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

func TestDataFrame_Concat(t *testing.T) {
	d1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewSeriesWithoutType("x", d1)
	df1 := NewDataFrame(s1)
	fmt.Println(df1)
}

func TestDataFrame_Concat1(t *testing.T) {
	d1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewSeriesWithoutType("x", d1)
	df1 := NewDataFrame(s1)
	d2 := []float64{101, 102}
	s2 := NewSeriesWithoutType("x", d2)
	df2 := NewDataFrame(s2)
	fmt.Println(df1)
	fmt.Println(df2)
	df3 := df1.Concat(df2)
	fmt.Println(df1)
	fmt.Println(df3)
}
