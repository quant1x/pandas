package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// FORCAST 返回S序列N周期回线性回归后的预测值
func FORCAST(S stat.Series, N any) any {
	return S.Rolling(N).Apply(func(X stat.Series, W num.DType) num.DType {
		x := X.DTypes()
		ws := num.Range[float64](int(W))
		c := num.PolyFit(ws, x, 1)
		w1 := num.Repeat[float64](W-1, int(W))
		vs := num.PolyVal(c, w1)

		if W > 1 {
			return vs[0]
		} else {
			return num.DTypeNaN
		}
	}).Values()
}
