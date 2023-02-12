package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Mean 求均值
func Mean[T Number](x []T) T {
	return unaryOperations1[T](x, vek32.Mean, vek.Mean, __mean_go[T])
}

func __mean_go[T Number](x []T) T {
	return __sum(x) / T(len(x))
}
