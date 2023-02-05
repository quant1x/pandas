package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// MA 计算移动均线
// 求序列的N日简单移动平均值, 返回序列
func MA(S pandas.Series, N any) any {
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
	return S.Rolling2(X).Mean().Values()
}
