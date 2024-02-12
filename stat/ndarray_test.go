package stat

import (
	"fmt"
	"gitee.com/quant1x/num"
	"reflect"
	"testing"
	"unsafe"
)

func TestNDArray_Len(t *testing.T) {
	f1 := []float64{1, 2, 3, 4, 5}
	a1 := NDArray[float64](f1)
	fmt.Println(a1)
	fmt.Println(a1.Len())
}

type X int

func TestNDArrayAll(t *testing.T) {
	var x1 X = 5
	var x2 int
	x2 = int(x1)
	fmt.Println(x2)
	d := []float32{1, 2, 3, 4, 5}
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	fmt.Printf("s : %#v\n", sh1)
	var s Series
	s = NDArray[float32](d)
	//s3 := []float32(s)
	//fmt.Println(s3)
	sh2 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("s : %#v\n", sh2.Data)
	fmt.Println(s.Len())
	s4 := s.Values()
	fmt.Println(s.Type())
	fmt.Println(s.Floats())

	f32 := ToFloat32(s)
	fmt.Println(f32)

	a1 := s.Diff(1)
	fmt.Println(a1)
	a2 := s.Ref(1)
	fmt.Println(a2)
	a2 = a2.FillNa(9, true)
	fmt.Println(a2)

	a3 := s.Mean()
	fmt.Println(a3)

	a4 := s.Shift(-1)
	fmt.Println(a4)
	s = s.Append(10, 11)
	fmt.Println(s)
	_ = s4
}

func TestNDArray_Rolling(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s := NewNDArray(d1...)
	r1 := s.Rolling(5).Mean()
	fmt.Println(r1)

	d2 := []float64{1, 2, 3, 4, 3, 3, 2, 1, num.Nil2Float64, num.Nil2Float64, num.Nil2Float64, num.Nil2Float64}
	r2 := s.Rolling(d2).Mean()
	fmt.Println(r2)
}

func TestNDArray_Apply(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDArray(d1...)
	fmt.Println(s1)
	s2 := s1.(NDArray[float64])
	s2.Apply2(func(idx int, v any) any {
		f := num.AnyToGeneric[float64](v)
		return f * f
	}, true)
	fmt.Println(s2)
}

func TestNDArray_Strings(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := NewNDArray[float64](d1...)
	ss := s1.Strings()
	fmt.Println(ss)
}

func TestNDArray_Bools(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 0, 7, 8, 9, 10, 11, 12}
	s1 := NewNDArray[float64](d1...)
	ss := s1.Bools()
	fmt.Println(ss)
}

func TestNDArray_Neq(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 0, 7, 8, 9, 10, 11, 12}
	s1 := NewNDArray[float64](d1...)
	ss := s1.Neq(0)
	fmt.Println(ss)
}

func TestNDArray_Not(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 0, 7, 8, 9, 10, 11, 12}
	s1 := NewNDArray[float64](d1...)
	ss := s1.Not()
	fmt.Println(ss)
}
