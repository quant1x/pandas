package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"unsafe"
)

// Repeat repeat
func Repeat[T StatType](f T, n int) []T {
	var d any
	bitSize := unsafe.Sizeof(f)
	if bitSize == 4 {
		d = vek32.Repeat(float32(f), n)
	} else if bitSize == 8 {
		d = vek.Repeat(float64(f), n)
	} else {
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
