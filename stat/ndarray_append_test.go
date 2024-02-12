package stat

import (
	"fmt"
	"testing"
)

func TestNDArray_Concat(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDArray(d1...)
	fmt.Println(s1)
	d2 := []float64{101, 102}
	s2 := NewNDArray(d2...)
	fmt.Println(s2)
	s3 := s1.Concat(s2)
	fmt.Println(s1)
	fmt.Println(s3)

}
