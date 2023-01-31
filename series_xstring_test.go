package pandas

import (
	"fmt"
	"testing"
)

func TestNewSeriesString(t *testing.T) {
	a3 := [3]string{"nan", "s1", "nan"}
	s3 := NewSeriesString("b", a3)
	fmt.Println(s3)

	var s1 Series
	s1 = NewSeriesString("a", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)
	expected := 4

	if s1.Len() != expected {
		t.Errorf("wrong val: expected: %v actual: %v", expected, s1.Len())
	}

	a2 := []string{"", "a", "b", "c"}
	s2 := NewSeriesString("b", a2)
	fmt.Println(s2)

	//s2 := s1.Shift(2)
	//fmt.Println(s2)
	//
	//s3 := s1.Repeat(1, 2)
	//fmt.Println(s3)
	//
	//s4 := NewSeriesInt64("x", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	//d4 := s4.Rolling(5).Mean()
	//fmt.Printf("d4 = %+v\n", d4.Values())
}
