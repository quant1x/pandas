package stat

import (
	"gitee.com/quant1x/vek"
	"gitee.com/quant1x/vek/vek32"
	"golang.org/x/exp/slices"
)

// Mul arithmetics 乘法
func Mul[T Number](x []T, y any) []T {
	return binaryOperations(x, y, vek32.Mul, vek.Mul, __mul_go[T])
}

func __mul_go[T Number](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] *= y[i]
	}
	return x
}
