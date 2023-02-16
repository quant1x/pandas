package v1

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestNewSeriesFloat64(t *testing.T) {
	var s1 Series
	s1 = NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)
	expected := 4

	if s1.Len() != expected {
		t.Errorf("wrong val: expected: %v actual: %v", expected, s1.Len())
	}
	s2 := s1.Shift(2)
	fmt.Println(s2.Values())

	s3 := s1.Repeat(1, 2)
	fmt.Println(s3.Values())

	s4 := NewSeriesFloat64("x", []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	d4 := s4.RollingV1(5).Mean()
	fmt.Printf("d4 = %+v\n", d4.Values())

	d5 := s4.RollingV1(5).StdDev()
	fmt.Printf("d5 = %+v\n", d5.Values())

	s5 := NewSeriesFloat64("x", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	e1 := s5.EWM(EW{Span: 5, Adjust: false}).Mean()
	fmt.Println(e1)

	e2 := s5.EWM(EW{Span: 5, Adjust: true}).Mean()
	fmt.Println(e2)

	e3 := s5.EWM(EW{Com: 5, Adjust: false}).Mean()
	fmt.Println(e3)

	e4 := s5.EWM(EW{Com: 5, Adjust: true}).Mean()
	fmt.Println(e4)

	e5 := s5.EWM(EW{Alpha: 1 / 5.0, Adjust: false}).Mean()
	fmt.Println(e5)
}

func TestSeriesWhere(t *testing.T) {
	c1 := []float64{1, 0, 3, 4, 5, 6, 7, 8, 9, 10}
	d1 := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	d2 := []float64{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}

	w1 := stat.Where(c1, d1, d2)
	fmt.Println(w1)

}