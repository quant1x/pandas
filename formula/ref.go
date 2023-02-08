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