package formula

import "gitee.com/quant1x/pandas/stat"

// STD 序列的N日标准差
func STD(S stat.Series, N any) stat.Series {
	return S.Rolling(N).Std()
}
