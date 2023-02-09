package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Min 计算最小值
func Min[T Float](f []T) T {
	if len(f) == 0 {
		return T(0)
	}
	var d any
	var s any
	s = f
	switch fs := s.(type) {
	case []float32:
		d = vek32.Min(fs)
	case []float64:
		d = vek.Min(fs)
	default:
		// 应该不会走到这里
		panic(ErrUnsupportedType)
	}
	return d.(T)
}
