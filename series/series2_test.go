package series

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type kSeries struct {
	columns []any
}

func TestSeries2_First(t *testing.T) {
	var columns []any
	columns = make([]any, 2)
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	columns[0] = x
	fmt.Printf("%+v\n", columns)
	x1 := make([]float64, 10)
	x2 := unsafe.Pointer(&x1)
	_ = x2
	s1 := NewSeries2("x", x)
	v1 := s1.First()
	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(v1)
	fmt.Printf("%+v\n", s1)
	s2 := s1.Shift(-5)
	fmt.Printf("%+v\n", s2)
}
