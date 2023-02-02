package pandas

import (
	"fmt"
	"testing"
)

func TestSeriesFrame(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := NewSeries(SERIES_TYPE_FLOAT, "x", data)
	fmt.Printf("%+v\n", s1)

	var d1 any
	d1 = data
	s2 := NewSeries(SERIES_TYPE_FLOAT, "x", d1)
	fmt.Printf("%+v\n", s2)

	var s3 Series
	s3 = NewSeriesBool("x", data)
	fmt.Printf("%+v\n", s3.Values())

	var s4 Series
	ts4 := GenericSeries[float64]("x", data...)
	s4 = *ts4
	fmt.Printf("%+v\n", s4.Values())
}

func TestNDFrameNew(t *testing.T) {
	// float64
	//d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, NaN(), 12}
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	nd1 := NewNDFrame[float64]("x", d1...)
	fmt.Println(nd1)
	fmt.Println(nd1.Records())
	nd11 := nd1.Subset(1, 2, true)
	fmt.Println(nd11.Records())
	fmt.Println(nd1.Max())
	fmt.Println(nd1.Rolling(5).Max())
	fmt.Println(nd1.Rolling(5).Min())

	nd12 := nd1.Rolling(5).Mean()
	d12 := nd12.Values()
	fmt.Println(d12)

	nd13 := nd1.Shift(3)
	fmt.Println(nd13.Values())
	nd14 := nd1.Rolling(5).StdDev()
	fmt.Println(nd14.Values())

	// string
	d2 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "nan", "12"}
	nd2 := NewNDFrame[string]("x", d2...)
	fmt.Println(nd2)
	nd21 := nd2.Rolling(5).Max()
	fmt.Println(nd21)
	nd2.FillNa(0, true)
	fmt.Println(nd2)
	fmt.Println(nd2.Records())
	fmt.Println(nd2.Empty())
}