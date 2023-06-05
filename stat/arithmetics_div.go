package stat

import (
	"gitee.com/quant1x/gox/num"
	"gitee.com/quant1x/gox/num/num32"
	"golang.org/x/exp/slices"
)

// Div arithmetics 除法
func Div[T Number](x []T, y any) []T {
	return binaryOperations(x, y, num32.Div, num.Div, __div_go[T])
}

func __div_go[T Number](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] /= y[i]
	}
	return x
}
