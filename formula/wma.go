package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// WMA 通达信S序列的N日加权移动平均 Yn = (1*X1+2*X2+3*X3+...+n*Xn)/(1+2+3+...+Xn)
func WMA(S pandas.Series, N any) any {
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
	return S.Rolling(X).Apply(func(S pandas.Series, N float32) float32 {
		if S.Len() == 0 {
			return stat.Nil2Float32
		}
		x := pandas.ToFloat32(S)
		//fmt.Println(x)
		x = stat.Reverse(x)
		//fmt.Println(x)
		v := stat.CumSum(x)
		//fmt.Println(v)
		v1 := stat.Sum(v)
		v2 := v1 * 2 / N / (N + 1)
		return v2
	}).Values()
}
