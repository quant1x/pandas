package formula

import (
	"gitee.com/quant1x/pandas"
)

// STD 序列的N日标准差
func STD(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Std()
}
