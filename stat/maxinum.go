package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"unsafe"
)

// Maxinum 两个序列横向比较最大值
func Maxinum[T Float](f1, f2 []T) []T {
	xlen := len(f1)
	ylen := len(f2)
	// 第找出最大长度

	maxLength := xlen
	if maxLength < ylen {
		maxLength = ylen
	}

	// 处理默认值
	defaultValue := typeDefault(T(0))
	// 对齐所有长度
	if xlen < maxLength {
		f1 = Align(f1, defaultValue, maxLength)
	}
	if ylen < maxLength {
		f2 = Align(f2, defaultValue, maxLength)
	}
	// 初始化返回值
	var s1, s2 any
	s1 = f1
	s2 = f2

	var d any

	bitSize := unsafe.Sizeof(f1[0])
	if bitSize == 4 {
		d = vek32.Maximum(s1.([]float32), s2.([]float32))
	} else if bitSize == 8 {
		d = vek.Maximum(s1.([]float64), s2.([]float64))
	} else {
		// 应该不会走到这里
		panic("不支持其它类型")
	}
	return d.([]T)
}
