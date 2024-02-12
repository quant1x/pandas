package formula

import (
	"gitee.com/quant1x/pandas"
)

// LLV 最近N周期的S最小值
func LLV(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Min()
}
