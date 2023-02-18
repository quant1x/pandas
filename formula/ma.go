package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// MA 计算移动均线
// 求序列的N日简单移动平均值, 返回序列
func MA(S stat.Series, N any) []stat.DType {
	return S.Rolling(N).Mean().DTypes()
}

func MA2(S stat.Series, N any) stat.Series {
	return S.Rolling(N).Mean()
}
