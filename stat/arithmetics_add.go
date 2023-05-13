package stat

import (
	"gitee.com/quant1x/gox/vek"
	"gitee.com/quant1x/gox/vek/vek32"
	"golang.org/x/exp/slices"
)

// Add arithmetics 加法
func Add[T Number](x []T, y any) []T {
	return binaryOperations(x, y, vek32.Add, vek.Add, __add_go[T])
}

func __add_go[T Number](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] += y[i]
	}
	return x
}
