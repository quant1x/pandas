package winpooh32

import "gitee.com/quant1x/pandas/winpooh32/math"

func fpEq(v1, v2, eps DType) bool {
	return math.Abs(v1-v2) < eps
}

func fpZero(v DType, eps DType) DType {
	switch {
	case math.Abs(v) < eps:
		return 0
	default:
		return v
	}
}

func IsNA(v DType) bool {
	return math.IsNaN(v) || math.IsInf(v, 0)
}
