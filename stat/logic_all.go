package stat

import "gitee.com/quant1x/vek"

// All 全部为真
func All[T Number | ~bool](x []T) bool {
	switch vs := any(x).(type) {
	case []bool:
		return vek.All(vs)
	case []int8:
		return __all_go(vs)
	case []uint8:
		return __all_go(vs)
	case []int16:
		return __all_go(vs)
	case []uint16:
		return __all_go(vs)
	case []int32:
		return __all_go(vs)
	case []uint32:
		return __all_go(vs)
	case []int64:
		return __all_go(vs)
	case []uint64:
		return __all_go(vs)
	case []int:
		return __all_go(vs)
	case []uint:
		return __all_go(vs)
	case []uintptr:
		return __all_go(vs)
	case []float32:
		return __all_go(vs)
	case []float64:
		return __all_go(vs)
	}
	return false
}

func __all_go[T Number](x []T) bool {
	for i := 0; i < len(x); i++ {
		if x[i] == 0 {
			return false
		}
	}
	return true
}
