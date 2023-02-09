package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"golang.org/x/exp/slices"
)

// CumSum 计算累和
func CumSum[T StatType](f []T) []T {
	if len(f) == 0 {
		return []T{}
	}
	var d any
	var s any
	s = f
	switch fs := s.(type) {
	case []float32:
		d = vek32.CumSum(fs)
	case []float64:
		d = vek.CumSum(fs)
	default:
		// 剩下的就是int32和int64, 循环吧
		sum := T(0)
		x := slices.Clone(f)
		for i := 0; i < len(x); i++ {
			sum += x[i]
			x[i] = sum
		}
		d = x
	}

	return d.([]T)
}
