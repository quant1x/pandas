package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// STD 序列的N日标准差
func STD(S pandas.Series, N any) pandas.Series {
	return v2STD(S, N)
}

func v1STD(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Std()
}

func v2STD(S pandas.Series, N any) pandas.Series {
	x := S.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, N, func(N num.DType, values ...int32) int32 {
			return num.Std(values)
		})
		return pandas.SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, N, func(N num.DType, values ...int64) int64 {
			return num.Std(values)
		})
		return pandas.SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, N, func(N num.DType, values ...float32) float32 {
			return num.Std(values)
		})
		return pandas.SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, N, func(N num.DType, values ...float64) float64 {
			return num.Std(values)
		})
		return pandas.SliceToSeries(d)
	}
	panic(num.ErrUnsupportedType)
}
