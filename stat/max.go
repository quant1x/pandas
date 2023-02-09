package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Max 计算最大值
func Max[T Float](f []T) T {
	if len(f) == 0 {
		return T(0)
	}

	var d any
	var s any
	s = f
	switch fs := s.(type) {
	case []float32:
		d = vek32.Max(fs)
	case []float64:
		d = vek.Max(fs)
	default:
		panic(ErrUnsupportedType)
	}

	return d.(T)
}
