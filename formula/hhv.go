package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// HHV 最近N周期的S最大值
func HHV(S stat.Series, N any) stat.Series {
	return S.Rolling(N).Max()
}
