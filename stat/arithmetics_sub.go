package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"golang.org/x/exp/slices"
)

// Sub arithmetics 减法
func Sub[T StatType](x []T, y any) []T {
	return binaryOperations(x, y, vek32.Sub, vek.Sub, __sub[T])
}

func __sub[T StatType](x, y []T) []T {
	x = slices.Clone(x)
	for i := 0; i < len(x); i++ {
		x[i] -= y[i]
	}
	return x
}
