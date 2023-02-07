package formula

import "gitee.com/quant1x/pandas"

// HHV 最近N周期的S最大值
func HHV(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Max()
}
