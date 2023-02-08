package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// ArgMin Returns the indices of the minimum values along an axis.
// 返回轴上最小值的索引
func ArgMin[T Number](v []T) int {
	var vv any = v
	switch values := vv.(type) {
	case []float32:
		return vek32.ArgMin(values)
	case []float64:
		return vek.ArgMin(values)
	default:
		return __arg_min(v)
	}

}

func __arg_min[T Number](x []T) int {
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
