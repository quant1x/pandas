package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// AVEDEV 平均绝对偏差, (序列与其平均值的绝对差的平均值)
//
//	AVEDEV(S,N) 返回平均绝对偏差
func AVEDEV(S pandas.Series, N any) any {
	return S.Rolling(N).Apply(func(X pandas.Series, W stat.DType) stat.DType {
		x := X.DTypes()
		x1 := X.Mean()
		r := stat.Sub(x, x1)
		r = stat.Abs(r)
		return stat.Mean(r)
	}).Values()
}