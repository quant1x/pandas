package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// ArgMax Returns the indices of the maximum values along an axis.
//
//	返回轴上最大值的索引
func ArgMax[T Number](x []T) int {
	ret := unaryOperations2[T, int](x, vek32.ArgMax, vek.ArgMax, __arg_max_go[T])
	return ret
}

func ArgMax2[T BaseType](x []T) int {
	var d int
	switch vs := any(x).(type) {
	case []float32:
		d = ArgMax(vs)
	case []float64:
		d = ArgMax(vs)
	case []int:
		d = ArgMax(vs)
	case []int8:
		d = ArgMax(vs)
	case []int16:
		d = ArgMax(vs)
	case []int32:
		d = ArgMax(vs)
	case []int64:
		d = ArgMax(vs)
	case []uint:
		d = ArgMax(vs)
	case []uint8:
		d = ArgMax(vs)
	case []uint16:
		d = ArgMax(vs)
	case []uint32:
		d = ArgMax(vs)
	case []uint64:
		d = ArgMax(vs)
	case []uintptr:
		d = ArgMax(vs)
	case []string:
		d = __arg_max_go(vs)
	case []bool:
		d = __arg_max_go_bool(vs)
	default:
		// 其它类型原样返回
		panic(Throw(any(x)))
	}

	return d
}

func __arg_max_go[T Ordered](x []T) int {
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

func __arg_max_go_bool(x []bool) int {
	max := BoolToInt(x[0])
	idx := 0
	for i, v := range x[1:] {
		if BoolToInt(v) > max {
			max = BoolToInt(v)
			idx = 1 + i
		}
	}
	return idx
}
