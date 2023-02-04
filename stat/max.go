package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"unsafe"
)

// Max 计算最大值
func Max[T Float](f []T) T {
	if len(f) == 0 {
		return T(0)
	}
	var d any
	var s any
	s = f
	bitSize := unsafe.Sizeof(f[0])
	if bitSize == 4 {
		d = vek32.Max(s.([]float32))
	} else if bitSize == 8 {
		d = vek.Max(s.([]float64))
	} else {
		// 应该不会走到这里
		d = T(0)
	}
	return d.(T)
}
