package stat

import (
	"gitee.com/quant1x/gox/num"
	"gitee.com/quant1x/gox/num/num32"
	"slices"
)

// Sub arithmetics 减法
func Sub[T Number](x []T, y any) []T {
	return binaryOperations(x, y, num32.Sub, num.Sub, __sub_go[T])
}

func __sub_go[T Number](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] -= y[i]
	}
	return x
}
