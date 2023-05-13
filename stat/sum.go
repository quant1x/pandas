package stat

import (
	"gitee.com/quant1x/gox/vek"
	"gitee.com/quant1x/gox/vek/vek32"
)

// Sum 求和
func Sum[T Number](f []T) T {
	if len(f) == 0 {
		return T(0)
	}
	var d any
	var s any
	s = f
	switch fs := s.(type) {
	case []float32:
		d = vek32.Sum(fs)
	case []float64:
		d = vek.Sum(fs)
	default:
		d = __sum(fs.([]T))
	}

	return d.(T)
}

func __sum[T Number](x []T) T {
	sum := T(0)
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
