package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"unsafe"
)

// Sum 计算累和
func Sum[T StatType](f []T) T {
	if len(f) == 0 {
		return T(0)
	}
	var d any
	var s any
	s = f
	bitSize := unsafe.Sizeof(f[0])
	if bitSize == 4 {
		d = vek32.Sum(s.([]float32))
	} else if bitSize == 8 {
		d = vek.Sum(s.([]float64))
	} else {
		// 剩下的就是int32和int64, 循环吧
		m := T(0)
		for _, v := range f {
			m += v
		}
		d = m
	}

	return d.(T)
}
