package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Repeat 构造n长度的f的泛型切片
func Repeat[T BaseType](f T, n int) []T {
	var d any
	var s any = f
	switch fs := s.(type) {
	case float32:
		d = vek32.Repeat(fs, n)
	case float64:
		d = vek.Repeat(fs, n)
	default:
		// 剩下非float32和float64, 循环吧
		d = []T{}
		m := make([]T, n)
		for i := 0; i < n; i++ {
			m[i] = f
		}
		d = m
	}
	return d.([]T)
}

// Range 产生从0到n-1的数组
func Range[T Number](n int) []T {
	var dest any

	var start T = 0
	var v any = start
	switch a := v.(type) {
	case float32:
		dest = vek32.Range(a, a+float32(n))
	case float64:
		dest = vek.Range(a, a+float64(n))
	default:
		// 其它类型
		d := make([]T, n)
		for i := 0; i < n; i++ {
			d[i] = start
			start += 1
		}
		dest = d
	}
	return dest.([]T)
}
