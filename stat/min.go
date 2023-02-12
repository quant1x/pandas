package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Min 纵向计算x最小值
func Min[T Number](x []T) T {
	return unaryOperations1[T](x, vek32.Min, vek.Min, __min_go[T])
}

func __min_go[T Number](x []T) T {
	min := x[0]
	for _, v := range x[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
