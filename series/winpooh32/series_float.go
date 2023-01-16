package winpooh32

import (
	"gitee.com/quant1x/pandas/series/winpooh32/math"
)

func fpEq(v1, v2, eps float64) bool {
	return math.Abs(v1-v2) < eps
}

func fpZero(v float64, eps float64) float64 {
	switch {
	case math.Abs(v) < eps:
		return 0
	default:
		return v
	}
}

func IsNA(v float64) bool {
	return math.IsNaN(v) || math.IsInf(v, 0)
}
