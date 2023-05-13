package stat

import (
	"gitee.com/quant1x/gox/vek"
	"gitee.com/quant1x/gox/vek/vek32"
)

// Mean 求均值
func Mean[T Number](x []T) T {
	return unaryOperations1[T](x, vek32.Mean, vek.Mean, __mean_go[T])
}

func __mean_go[T Number](x []T) T {
	return __sum(x) / T(len(x))
}

func Mean2[T BaseType](x []T) T {
	var d any
	switch vs := any(x).(type) {
	case []float32:
		d = Mean(vs)
	case []float64:
		d = Mean(vs)
	case []int:
		d = Mean(vs)
	case []int8:
		d = Mean(vs)
	case []int16:
		d = Mean(vs)
	case []int32:
		d = Mean(vs)
	case []int64:
		d = Mean(vs)
	case []uint:
		d = Mean(vs)
	case []uint8:
		d = Mean(vs)
	case []uint16:
		d = Mean(vs)
	case []uint32:
		d = Mean(vs)
	case []uint64:
		d = Mean(vs)
	case []uintptr:
		d = Mean(vs)
	//case []string:
	//	d = __max_go(vs)
	default:
		// 其它类型原样返回
		panic(Throw(any(x)))
	}

	return d.(T)
}
