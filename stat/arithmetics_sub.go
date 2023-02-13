package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"golang.org/x/exp/slices"
)

// Sub arithmetics 减法
func Sub[T Number](x []T, y any) []T {
	return binaryOperations(x, y, vek32.Sub, vek.Sub, __sub_go[T])
}

func __sub_go[T Number](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] -= y[i]
	}
	return x
}
