package formula

import (
	"gitee.com/quant1x/pandas"
)

// SUM 求累和
// 如果N=0, 则从第一个有效值累加到当前
// 下一步再统一返回值
func SUM(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Sum()
}
