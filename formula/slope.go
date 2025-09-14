package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// SLOPE 计算周期回线性回归斜率
//
//	SLOPE(S,N) 返回线性回归斜率,N支持变量
func SLOPE(S pandas.Series, N any) any {
	return S.Rolling(N).Apply(func(X pandas.Series, W num.DType) num.DType {
		x := X.DTypes()
		w := num.Range[num.DType](len(x))
		c := num.PolyFit(w, x, 1)
		return c[0]
	}).Values()
}
