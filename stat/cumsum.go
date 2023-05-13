package stat

import (
	"gitee.com/quant1x/gox/vek"
	"gitee.com/quant1x/gox/vek/vek32"
	"golang.org/x/exp/slices"
)

// CumSum 计算累和
func CumSum[T Number](x []T) []T {
	return unaryOperations(x, vek32.CumSum, vek.CumSum, __cumsum_go[T])
}

//func CumSum[T StatType](x []T) []T {
//	if len(x) == 0 {
//		return []T{}
//	}
//	var d any
//	var s any
//	s = x
//	switch fs := s.(type) {
//	case []float32:
//		d = vek32.CumSum(fs)
//	case []float64:
//		d = vek.CumSum(fs)
//	default:
//		// 剩下的就是int32和int64, 循环吧
//		sum := T(0)
//		x := slices.Clone(f)
//		for i := 0; i < len(x); i++ {
//			sum += x[i]
//			x[i] = sum
//		}
//		d = x
//	}
//
//	return d.([]T)
//}

func __cumsum_go[T Number](x []T) []T {
	x = slices.Clone(x)
	sum := T(0)
	for i := 0; i < len(x); i++ {
		sum += x[i]
		x[i] = sum
	}
	return x
}
