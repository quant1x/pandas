package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Repeat repeat
func Repeat[T StatType](f T, n int) []T {
	var d any
	var s any = f
	switch fs := s.(type) {
	case float32:
		d = vek32.Repeat(fs, n)
	case float64:
		d = vek.Repeat(fs, n)
	default:
		// 应该不会走到这里
		d = []T{}
		// 剩下的就是int32和int64, 循环吧
		m := make([]T, n)
		for i := 0; i < n; i++ {
			m[i] = f
		}
		d = m
	}
	return d.([]T)
}
