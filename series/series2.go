package series

import (
	"gitee.com/quant1x/pandas/algorithms/winpooh32/math"
	"reflect"
)

//type S[T int64 | float64] []T

type Series2[T ~int64 | ~float64] struct {
	data []T          // The values of the elements
	Name string       // The name of the series
	t    reflect.Type // The type of the series
}

func NewSeries2[T ~int64 | ~float64](name string, data []T) Series2[T] {
	val := data[0]
	s := Series2[T]{
		Name: name,
		data: data,
		t:    reflect.TypeOf(val),
	}

	return s
}

func (s Series2[T]) First() T {
	values := s.data
	return values[0]
}

// Shift shifts values by specified periods count.
func (s Series2[T]) Shift(periods int) Series2[T] {
	if periods == 0 {
		return s
	}

	values := s.data

	var (
		naVals []T
		dst    []T
		src    []T
	)
	if shlen := int(math.Abs(float64(periods))); shlen < len(values) {
		if periods > 0 {
			naVals = values[:shlen]
			dst = values[shlen:]
			src = values
		} else {
			naVals = values[len(values)-shlen:]
			dst = values[:len(values)-shlen]
			src = values[shlen:]
		}

		copy(dst, src)
	} else {
		naVals = values
	}

	for i := range naVals {
		switch s.t.Kind() {
		case reflect.Int64, reflect.Float64:
			naVals[i] = T(math.NaN())
		}

	}

	return s
}
