package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Sum 求和
func Sum[T StatType](f []T) T {
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
		// 剩下的就是int32和int64, 循环吧
		m := T(0)
		for _, v := range f {
			m += v
		}
		d = m
	}

	return d.(T)
}
