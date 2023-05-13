package stat

import (
	"gitee.com/quant1x/gox/vek"
	"gitee.com/quant1x/gox/vek/vek32"
	"golang.org/x/exp/slices"
)

// Minimum AVX2版本, 两个序列横向比较最大值
func Minimum[T Number](f1, f2 []T) []T {
	xlen := len(f1)
	ylen := len(f2)

	// 第找出最大长度
	maxLength := xlen
	if maxLength < ylen {
		maxLength = ylen
	}

	// 处理默认值
	defaultValue := typeDefault[T]()
	// 对齐所有长度
	if xlen < maxLength {
		f1 = Align(f1, T(defaultValue), maxLength)
	}
	if ylen < maxLength {
		f2 = Align(f2, T(defaultValue), maxLength)
	}
	// 初始化返回值
	var s1, s2 any
	s1 = f1
	s2 = f2

	var d any
	switch fs1 := s1.(type) {
	case []float32:
		d = vek32.Minimum(fs1, s2.([]float32))
	case []float64:
		d = vek.Minimum(fs1, s2.([]float64))
	default:
		// 目前暂时走不到这里
		f1 = slices.Clone(f1)
		__minimum_go(f1, f2)
		d = f1
	}
	return d.([]T)
}

func __minimum_go[T Number](x, y []T) {
	for i := 0; i < len(x); i++ {
		if y[i] < x[i] {
			x[i] = y[i]
		}
	}
}
