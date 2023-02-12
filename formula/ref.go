package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// REF 引用前N的序列
func REF(S pandas.Series, N any) any {
	var X []float32
	switch v := N.(type) {
	case int:
		X = stat.Repeat[float32](float32(v), S.Len())
	case pandas.Series:
		vs := v.Values()
		X = pandas.SliceToFloat32(vs)
		X = stat.Align(X, pandas.Nil2Float32, S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	return S.Ref(X).Values()
}

func REF2[T stat.GenericType](S []T, N any) []T {
	sLen := len(S)
	var X []stat.DType
	switch v := N.(type) {
	case int:
		X = stat.Repeat[stat.DType](stat.DType(v), sLen)
	case []stat.DType:
		X = stat.Align(X, stat.Nil2Float64, sLen)
	default:
		panic(exception.New(1, "error window"))
	}
	return stat.Shift2[T](S, X)
}
