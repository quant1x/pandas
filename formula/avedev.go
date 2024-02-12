package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// AVEDEV 平均绝对偏差, (序列与其平均值的绝对差的平均值)
//
//	AVEDEV(S,N) 返回平均绝对偏差
func AVEDEV(S stat.Series, N any) stat.Series {
	return S.Rolling(N).Apply(func(X stat.Series, W num.DType) num.DType {
		x := X.DTypes()
		x1 := X.Mean()
		r := num.Sub(x, x1)
		r = num.Abs(r)
		return num.Mean(r)
	})
}
