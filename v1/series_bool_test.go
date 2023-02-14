package v1

import (
	"fmt"
	"testing"
)

func TestNewSeriesBool(t *testing.T) {
	var s1 Series
	s1 = NewSeriesBool("sales", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)
	expected := 4

	if s1.Len() != expected {
		t.Errorf("wrong val: expected: %v actual: %v", expected, s1.Len())
	}
	s11 := s1.Copy()
	fmt.Println(s11.Values())
	s2 := s1.Shift(2)
	fmt.Println(s2.Values())

	s3 := s1.Repeat(1, 2)
	fmt.Println(s3.Values())

	//s4 := NewSeriesBool("x", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	//d4 := s4.RollingV1(5).Mean()
	//fmt.Printf("d4 = %+v\n", d4.Values())
}
