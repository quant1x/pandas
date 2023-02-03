package avx2

import "github.com/viterin/vek"

func ToBool(x []float64) []bool {
	return vek.ToBool(x)
}
