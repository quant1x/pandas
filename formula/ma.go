package formula

import (
	"gitee.com/quant1x/pandas"
)

// MA 计算移动均线
//
//	求序列的N日简单移动平均值, 返回序列
func MA(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Mean()
}
