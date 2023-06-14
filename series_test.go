package pandas

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestSeriesFrame(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := NewSeries(stat.SERIES_TYPE_FLOAT64, "x", data)
	fmt.Printf("%+v\n", s1)

	var d1 any
	d1 = data
	s2 := NewSeries(stat.SERIES_TYPE_FLOAT64, "x", d1)
	fmt.Printf("%+v\n", s2)

	var s3 stat.Series
	// s3 = NewSeriesBool("x", data)
	s3 = NewSeries(stat.SERIES_TYPE_BOOL, "x", data)
	fmt.Printf("%+v\n", s3.Values())

	var s4 stat.Series
	ts4 := GenericSeries[float64]("x", data...)
	ts4 = NewSeries(stat.SERIES_TYPE_FLOAT64, "x", data)
	s4 = ts4
	fmt.Printf("%+v\n", s4.Values())
}

func TestNDFrameNew(t *testing.T) {
	// float64
	//d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, NaN(), 12}
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	nd1 := NewNDFrame[float64]("x", d1...)
	fmt.Println(nd1)

	r := api.RangeFinite(-1)
	ndr1 := nd1.Select(r)
	fmt.Println(ndr1.Values())

	fmt.Println(nd1.Records())
	nd11 := nd1.Subset(1, 2, true)
	fmt.Println(nd11.Records())
	fmt.Println(nd1.Max())

	nd13 := nd1.Shift(3)
	fmt.Println(nd13.Values())

	// string
	d2 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "nan", "12"}
	nd2 := NewNDFrame[string]("x", d2...)
	fmt.Println(nd2)
	nd2.FillNa(0, true)
	fmt.Println(nd2)
	fmt.Println(nd2.Records())
	fmt.Println(nd2.Empty())
}

func TestRolling2(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDFrame[float64]("x", d1...)
	df := NewDataFrame(s1)
	fmt.Println(df)
	fmt.Println("------------------------------------------------------------")

	N := 5
	fmt.Println("固定的参数, N =", N)
	r1 := df.Col("x").Rolling(5).Mean().Values()
	fmt.Println("序列化结果:", r1)
	fmt.Println("------------------------------------------------------------")
	d2 := []float64{1, 2, 3, 4, 3, 3, 2, 1, stat.Nil2Float64, stat.Nil2Float64, stat.Nil2Float64, stat.Nil2Float64}
	s2 := NewSeries(stat.SERIES_TYPE_FLOAT64, "x", d2)
	fmt.Printf("序列化参数: %+v\n", s2.Values())
	r2 := df.Col("x").Rolling(s2).Mean().Values()
	fmt.Println("序列化结果:", r2)
}

func TestSeriesConcat(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDFrame[float64]("x", d1...)
	d2 := []float64{101, 102}
	s2 := NewNDFrame[float64]("x", d2...)
	s21 := s2.Apply2(func(idx int, v any) any {
		r := v.(float64)
		return r * r
	}, true)
	fmt.Println(s21)
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := s1.Concat(s2)
	fmt.Println(s1)
	fmt.Println(s3)
}

func TestNDFrame_Strings(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDFrame[float64]("x", d1...)
	ss := s1.Strings()
	fmt.Println(ss)
}

func TestNDFrame_Bools(t *testing.T) {
	d1 := []float64{1, 2, 0, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDFrame[float64]("x", d1...)
	ss := s1.Bools()
	fmt.Println(ss)
}
