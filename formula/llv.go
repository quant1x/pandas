package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// LLV 最近N周期的S最小值
func LLV(S stat.Series, N any) stat.Series {
	return S.Rolling(N).Min()
}
