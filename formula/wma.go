package formula

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// WMA 通达信S序列的N日加权移动平均 Yn = (1*X1+2*X2+3*X3+...+n*Xn)/(1+2+3+...+Xn)
func WMA(S stat.Series, N any) stat.Series {
	var X []num.DType
	switch v := N.(type) {
	case int:
		X = num.Repeat[num.DType](num.DType(v), S.Len())
	case stat.Series:
		vs := v.DTypes()
		X = num.Align(vs, num.DTypeNaN, S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	d := S.Rolling(X).Apply(func(S stat.Series, N num.DType) num.DType {
		if S.Len() == 0 {
			return num.DTypeNaN
		}
		x := S.DTypes()
		x = api.Reverse(x)
		v := num.CumSum(x)
		v1 := num.Sum(v)
		v2 := v1 * 2 / N / (N + 1)
		if num.DTypeIsNaN(v2) {
			v2 = num.DTypeNaN
		}
		return v2
	})
	return d
}
