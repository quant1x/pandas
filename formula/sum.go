package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// SUM 求累和
// 如果N=0, 则从第一个有效值累加到当前
// 下一步再统一返回值
func SUM(S pandas.Series, N any) pandas.Series {
	return v2SUM(S, N)
}

func v1SUM(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Sum()
}

func v2SUM(S pandas.Series, N any) pandas.Series {
	x := S.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, N, func(N num.DType, values ...int32) int32 {
			return num.Sum(values)
		})
		return pandas.SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, N, func(N num.DType, values ...int64) int64 {
			return num.Sum(values)
		})
		return pandas.SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, N, func(N num.DType, values ...float32) float32 {
			return num.Sum(values)
		})
		return pandas.SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, N, func(N num.DType, values ...float64) float64 {
			return num.Sum(values)
		})
		return pandas.SliceToSeries(d)
	}
	panic(num.ErrUnsupportedType)
}
