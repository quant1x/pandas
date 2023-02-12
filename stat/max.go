package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Max 纵向计算x最大值
func Max[T Number](x []T) T {
	return unaryOperations1[T](x, vek32.Max, vek.Max, __max_go[T])
}

//func Max[T Float](f []T) T {
//	if len(f) == 0 {
//		return T(0)
//	}
//
//	var d any
//	var s any
//	s = f
//	switch fs := s.(type) {
//	case []float32:
//		d = vek32.Max(fs)
//	case []float64:
//		d = vek.Max(fs)
//	default:
//		panic(ErrUnsupportedType)
//	}
//
//	return d.(T)
//}

func __max_go[T Number](x []T) T {
	max := x[0]
	for _, v := range x[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
