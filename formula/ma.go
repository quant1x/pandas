package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// MAV1 计算移动均线
//
//	求序列的N日简单移动平均值, 返回序列
//	Deprecated: 推荐 MA
func MAV1(S pandas.Series, N any) []num.DType {
	return S.Rolling(N).Mean().DTypes()
}

func MA(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Mean()
}

// MAx 增量MA
// Deprecated: 错误的 [wangfeng on 2024/2/11 16:44]
func MAx(n int, old, new num.DType) num.DType {
	if n < 1 {
		return new
	}
	x := old * num.DType(n-1)
	x = (x + new) / num.DType(n)
	return x
}
