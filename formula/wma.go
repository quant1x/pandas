package formula

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// WMA 通达信S序列的N日加权移动平均 Yn = (1*X1+2*X2+3*X3+...+n*Xn)/(1+2+3+...+Xn)
func WMA(S pandas.Series, N any) pandas.Series {
	var X []num.DType
	switch v := N.(type) {
	case int:
		X = num.Repeat[num.DType](num.DType(v), S.Len())
	case pandas.Series:
		vs := v.DTypes()
		X = num.Align(vs, num.NaN(), S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	d := S.Rolling(X).Apply(func(S pandas.Series, N num.DType) num.DType {
		if S.Len() == 0 {
			return num.NaN()
		}
		x := S.DTypes()
		x = api.Reverse(x)
		v := num.CumSum(x)
		v1 := num.Sum(v)
		v2 := v1 * 2 / N / (N + 1)
		if num.DTypeIsNaN(v2) {
			v2 = num.NaN()
		}
		return v2
	})
	return d
}
