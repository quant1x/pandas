package stat

import "gitee.com/quant1x/vek"

// Any 任意一个为真
func Any[T Number | ~bool](x []T) bool {
	switch vs := any(x).(type) {
	case []bool:
		return vek.Any(vs)
	case []int8:
		return __any_go(vs)
	case []uint8:
		return __any_go(vs)
	case []int16:
		return __any_go(vs)
	case []uint16:
		return __any_go(vs)
	case []int32:
		return __any_go(vs)
	case []uint32:
		return __any_go(vs)
	case []int64:
		return __any_go(vs)
	case []uint64:
		return __any_go(vs)
	case []int:
		return __any_go(vs)
	case []uint:
		return __any_go(vs)
	case []uintptr:
		return __any_go(vs)
	case []float32:
		return __any_go(vs)
	case []float64:
		return __any_go(vs)
	}
	return false
}

func __any_go[T Number](x []T) bool {
	for i := 0; i < len(x); i++ {
		if x[i] != 0 {
			return true
		}
	}
	return false
}
