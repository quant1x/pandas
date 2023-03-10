package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// FORCAST 返回S序列N周期回线性回归后的预测值
func FORCAST(S stat.Series, N any) any {
	return S.Rolling(N).Apply(func(X stat.Series, W stat.DType) stat.DType {
		x := X.DTypes()
		ws := stat.Range[float64](int(W))
		c := stat.PolyFit(ws, x, 1)
		w1 := stat.Repeat[float64](W-1, int(W))
		vs := stat.PolyVal(c, w1)

		if W > 1 {
			return vs[0]
		} else {
			return stat.DTypeNaN
		}
	}).Values()
}
