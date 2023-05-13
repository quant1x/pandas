package stat

import "gitee.com/quant1x/gox/vek"

// Count 统计
func Count[T Number | ~bool](x []T) int {
	switch vs := any(x).(type) {
	case []bool:
		return vek.Count(vs)
	case []int8:
		return __count_go(vs)
	case []uint8:
		return __count_go(vs)
	case []int16:
		return __count_go(vs)
	case []uint16:
		return __count_go(vs)
	case []int32:
		return __count_go(vs)
	case []uint32:
		return __count_go(vs)
	case []int64:
		return __count_go(vs)
	case []uint64:
		return __count_go(vs)
	case []int:
		return __count_go(vs)
	case []uint:
		return __count_go(vs)
	case []uintptr:
		return __count_go(vs)
	case []float32:
		return __count_go(vs)
	case []float64:
		return __count_go(vs)
	}
	return 0
}

func __count_go[T Number](x []T) int {
	cnt := 0
	for i := 0; i < len(x); i++ {
		if x[i] != 0 {
			cnt += 1
		}
	}
	return cnt
}
