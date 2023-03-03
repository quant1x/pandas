package formula

import (
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/exception"
)

// WMA 通达信S序列的N日加权移动平均 Yn = (1*X1+2*X2+3*X3+...+n*Xn)/(1+2+3+...+Xn)
func WMA(S stat.Series, N any) stat.Series {
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
	d := S.Rolling(X).Apply(func(S stat.Series, N stat.DType) stat.DType {
		if S.Len() == 0 {
			return stat.DTypeNaN
		}
		x := S.DTypes()
		x = stat.Reverse(x)
		v := stat.CumSum(x)
		v1 := stat.Sum(v)
		v2 := v1 * 2 / N / (N + 1)
		return v2
	})
	return d
}
