package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"golang.org/x/exp/slices"
)

// Div arithmetics 除法
func Div[T Number](x []T, y any) []T {
	return binaryOperations(x, y, vek32.Div, vek.Div, __div_go[T])
}

func __div_go[T Number](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] /= y[i]
	}
	return x
}
