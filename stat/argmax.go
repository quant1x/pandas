package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// ArgMax Returns the indices of the maximum values along an axis.
// 返回轴上最大值的索引
//func ArgMax_V1[T Number](v []T) int {
//	var vv any = v
//	switch values := vv.(type) {
//	case []float32:
//		return vek32.ArgMax(values)
//	case []float64:
//		return vek.ArgMax(values)
//	default:
//		return __arg_max(v)
//	}
//}

func __arg_max_go[T Number](x []T) int {
	max := x[0]
	idx := 0
	for i, v := range x[1:] {
		if v > max {
			max = v
			idx = 1 + i
		}
	}
	return idx
}

// ArgMax Returns the indices of the maximum values along an axis.
//
//	返回轴上最大值的索引
func ArgMax[T Number](x []T) int {
	ret := unaryOperations2[T, int](x, vek32.ArgMax, vek.ArgMax, __arg_max_go[T])
	return ret
}
