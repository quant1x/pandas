package formula

import (
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// REF 引用前N的序列
func REF(S stat.Series, N any) stat.Series {
	var X []stat.DType
	switch v := N.(type) {
	case int:
		X = stat.Repeat[stat.DType](stat.DType(v), S.Len())
	case stat.Series:
		vs := v.DTypes()
		X = stat.Align(vs, stat.DTypeNaN, S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	return S.Ref(X)
}

func REF2[T stat.GenericType](S []T, N any) []T {
	sLen := len(S)
	var X []stat.DType
	switch v := N.(type) {
	case int:
		X = stat.Repeat[stat.DType](stat.DType(v), sLen)
	case []stat.DType:
		X = stat.Align(v, stat.DTypeNaN, sLen)
	default:
		panic(exception.New(1, "error window"))
	}
	return stat.Shift2[T](S, X)
}
