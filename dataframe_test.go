package pandas

import (
	"fmt"
	"testing"
)

func TestDataFrameT0(t *testing.T) {
	var s1 Series
	s1 = NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)
	expected := 4

	if s1.Len() != expected {
		t.Errorf("wrong val: expected: %v actual: %v", expected, s1.Len())
	}
	s2 := s1.Shift(-2)
	df := NewDataFrame(s1, s2)
	fmt.Println(df)
	df.FillNa(0.00, true)
	fmt.Println(df)

	_ = s2
}

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

func TestDataFrame_Join(t *testing.T) {
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
	s_e := GenericSeries[string]("", "a0", "a1", "a2", "a3")
	df2 := df1.Join(s_e)
	fmt.Println(df2)
}
