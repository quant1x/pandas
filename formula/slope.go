package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// SLOPE 计算周期回线性回归斜率
//
//	SLOPE(S,N) 返回线性回归斜率,N支持变量
func SLOPE(S pandas.Series, N any) any {
	return S.Rolling(N).Apply(func(X pandas.Series, W stat.DType) stat.DType {
		x := X.DTypes()
		w := stat.Sequence[stat.DType](len(x))
		c := stat.PolyFit(w, x, 1)
		return c[0]
	}).Values()
}
