package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// LLV 最近N周期的S最小值
func LLV(S pandas.Series, N any) pandas.Series {
	return v2LLV(S, N)
}

func v1LLV(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Min()
}

func v2LLV(S pandas.Series, N any) pandas.Series {
	x := S.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, N, func(N num.DType, values ...int32) int32 {
			return num.Min2(values)
		})
		return pandas.SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, N, func(N num.DType, values ...int64) int64 {
			return num.Min2(values)
		})
		return pandas.SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, N, func(N num.DType, values ...float32) float32 {
			return num.Min2(values)
		})
		return pandas.SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, N, func(N num.DType, values ...float64) float64 {
			return num.Min2(values)
		})
		return pandas.SliceToSeries(d)
	case []string:
		d := num.RollingV1(vs, N, func(N num.DType, values ...string) string {
			return num.Min2(values)
		})
		return pandas.SliceToSeries(d)
	}
	panic(num.ErrUnsupportedType)
}
