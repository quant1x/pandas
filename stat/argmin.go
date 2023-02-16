package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// ArgMin Returns the indices of the minimum values along an axis.
//
//	返回轴上最小值的索引
func ArgMin[T Number](x []T) int {
	ret := unaryOperations2[T, int](x, vek32.ArgMin, vek.ArgMin, __arg_min_go[T])
	return ret
}

func ArgMin2[T BaseType](x []T) int {
	var d int
	switch vs := any(x).(type) {
	case []float32:
		d = ArgMin(vs)
	case []float64:
		d = ArgMin(vs)
	case []int:
		d = ArgMin(vs)
	case []int8:
		d = ArgMin(vs)
	case []int16:
		d = ArgMin(vs)
	case []int32:
		d = ArgMin(vs)
	case []int64:
		d = ArgMin(vs)
	case []uint:
		d = ArgMin(vs)
	case []uint8:
		d = ArgMin(vs)
	case []uint16:
		d = ArgMin(vs)
	case []uint32:
		d = ArgMin(vs)
	case []uint64:
		d = ArgMin(vs)
	case []uintptr:
		d = ArgMin(vs)
	case []string:
		d = __arg_min_go(vs)
	case []bool:
		d = __arg_min_go_bool(vs)
	default:
		// 其它类型原样返回
		panic(Throw(any(x)))
	}

	return d
}

func __arg_min_go[T Ordered](x []T) int {
	min := x[0]
	idx := 0
	for i, v := range x[1:] {
		if v < min {
			min = v
			idx = 1 + i
		}
	}
	return idx
}

func __arg_min_go_bool(x []bool) int {
	min := bool2Int(x[0])
	idx := 0
	for i, v := range x[1:] {
		if bool2Int(v) < min {
			min = bool2Int(v)
			idx = 1 + i
		}
	}
	return idx
}
