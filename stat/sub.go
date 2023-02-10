package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"golang.org/x/exp/slices"
)

// Sub 减法
func Sub[T StatType](x []T, y any) []T {
	var d any
	xLen := len(x)
	var s any = x
	switch vs := s.(type) {
	case []float32:
		f32s := AnyToSlice[float32](y, xLen)
		d = vek32.Sub(vs, f32s)
	case []float64:
		f64s := AnyToSlice[float64](y, xLen)
		d = vek.Sub(vs, f64s)
	default:
		ts := AnyToSlice[T](y, xLen)
		d = __sub(x, ts)
	}
	return d.([]T)
}

func __sub[T StatType](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] -= y[i]
	}
	return x
}
