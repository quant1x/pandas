package stat

import "github.com/viterin/vek"

// And 两者为真
func And[T Number | ~bool](x, y []T) []bool {
	switch vs := any(x).(type) {
	case []bool:
		return vek.And(vs, any(y).([]bool))
	case []int8:
		return __and_go(vs, any(y).([]int8))
	case []uint8:
		return __and_go(vs, any(y).([]uint8))
	case []int16:
		return __and_go(vs, any(y).([]int16))
	case []uint16:
		return __and_go(vs, any(y).([]uint16))
	case []int32:
		return __and_go(vs, any(y).([]int32))
	case []uint32:
		return __and_go(vs, any(y).([]uint32))
	case []int64:
		return __and_go(vs, any(y).([]int64))
	case []uint64:
		return __and_go(vs, any(y).([]uint64))
	case []int:
		return __and_go(vs, any(y).([]int))
	case []uint:
		return __and_go(vs, any(y).([]uint))
	case []uintptr:
		return __and_go(vs, any(y).([]uintptr))
	case []float32:
		return __and_go(vs, any(y).([]float32))
	case []float64:
		return __and_go(vs, any(y).([]float64))
	}
	panic(Throw(x))
}

func __and_go[T Number](x, y []T) []bool {
	d := make([]bool, len(x))
	for i := 0; i < len(x); i++ {
		if x[i] != 0 && y[i] != 0 {
			d[i] = true
		} else {
			d[i] = false
		}
	}
	return d
}
