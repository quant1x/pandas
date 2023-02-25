package stat

import (
	"math"
)

func __min_n_go[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func __float_exp_go(f float64) (frac float64, exp int) {
	return math.Frexp(f)
}
