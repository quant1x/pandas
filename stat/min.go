package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"unsafe"
)

// Min 计算最小值
func Min[T Float](f []T) T {
	if len(f) == 0 {
		return T(0)
	}
	var d any
	var s any
	s = f
	bitSize := unsafe.Sizeof(f[0])
	if bitSize == 4 {
		d = vek32.Min(s.([]float32))
	} else if bitSize == 8 {
		d = vek.Min(s.([]float64))
	} else {
		// 应该不会走到这里
		d = T(0)
	}
	return d.(T)
}