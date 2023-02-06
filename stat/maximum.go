package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"unsafe"
)

// MaximumAvx2 两个序列横向比较最大值
// TODO:print(np.maximum(1.4, np.nan)) 输出nan
func MaximumAvx2[T Float](f1, f2 []T) []T {
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
		panic("other types are not supported")
	}
	return d.([]T)
}

// Maximum 两个序列横向比较最大值
// TODO:print(np.maximum(1.4, np.nan)) 输出nan
func Maximum[T Float](f1, f2 []T) []T {
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
	d := make([]T, maxLength)
	bitSize := unsafe.Sizeof(f1[0])
	for i := 0; i < maxLength; i++ {
		if Float64IsNaN(float64(f1[i])) || Float64IsNaN(float64(f2[i])) {
			if bitSize == 4 {
				d[i] = T(Nil2Float32)
			} else {
				d[i] = T(Nil2Float64)
			}
		}
		if f1[i] > f2[i] {
			d[i] = f1[i]
		} else {
			d[i] = f2[i]
		}
	}
	return d
}
