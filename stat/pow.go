package stat

import (
	"math"
)

func Pow[T Number](v []T, n int) []T {
	x := make([]T, len(v))
	for i := 0; i < len(v); i++ {
		x[i] = __pow_go(v[i], n)
	}
	return x
}

func __pow_go[T Number](x T, n int) T {
	y := math.Pow(float64(x), float64(n))
	return T(y)
}
