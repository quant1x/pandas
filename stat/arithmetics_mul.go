package stat

import (
	"gitee.com/quant1x/gox/num"
	"gitee.com/quant1x/gox/num/num32"
	"golang.org/x/exp/slices"
)

// Mul arithmetics 乘法
func Mul[T Number](x []T, y any) []T {
	return binaryOperations(x, y, num32.Mul, num.Mul, __mul_go[T])
}

func __mul_go[T Number](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] *= y[i]
	}
	return x
}
