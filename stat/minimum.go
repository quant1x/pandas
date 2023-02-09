package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"golang.org/x/exp/slices"
)

// Minimum AVX2版本, 两个序列横向比较最大值
func Minimum[T Float](f1, f2 []T) []T {
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
	switch fs1 := s1.(type) {
	case []float32:
		d = vek32.Minimum(fs1, s2.([]float32))
	case []float64:
		d = vek.Minimum(fs1, s2.([]float64))
	default:
		// 目前暂时走不到这里
		f1 = slices.Clone(f1)
		__minimum(f1, f2)
		d = f1
	}
	return d.([]T)
}

// Minimum_GO go版本 两个序列横向比较最大值
func Minimum_GO[T Float](f1, f2 []T) []T {
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
	for i := 0; i < maxLength; i++ {
		if Float64IsNaN(float64(f1[i])) || Float64IsNaN(float64(f2[i])) {
			var s1 any = f1[i]
			switch s1.(type) {
			case float32:
				d[i] = T(Nil2Float32)
			case []float64:
				d[i] = T(Nil2Float64)
			default:
				panic(ErrUnsupportedType)
			}
			continue
		}
		if f1[i] < f2[i] {
			d[i] = f1[i]
		} else {
			d[i] = f2[i]
		}
	}
	return d
}

// 暂时用不到, 先放在这里, 以后可能要扩展类型
func __minimum[T Float](x, y []T) {
	for i := 0; i < len(x); i++ {
		if y[i] < x[i] {
			x[i] = y[i]
		}
	}
}
