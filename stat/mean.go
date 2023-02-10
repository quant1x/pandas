package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Mean 求均值
func Mean[T StatType](x []T) T {
	var d any
	var s any = x
	switch vs := s.(type) {
	case []float32:
		d = vek32.Mean(vs)
	case []float64:
		d = vek.Mean(vs)
	default:
		d = __mean(x)
	}
	return d.(T)
}

func __mean[T StatType](x []T) T {
	return __sum(x) / T(len(x))
}
