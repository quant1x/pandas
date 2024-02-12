package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// AVEDEV 平均绝对偏差, (序列与其平均值的绝对差的平均值)
//
//	AVEDEV(S,N) 返回平均绝对偏差
func AVEDEV(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Apply(func(X pandas.Series, W num.DType) num.DType {
		x := X.DTypes()
		x1 := X.Mean()
		r := num.Sub(x, x1)
		r = num.Abs(r)
		return num.Mean(r)
	})
}
