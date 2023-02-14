package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// SQRT 求S的平方根
func SQRT(S stat.Series) []stat.DType {
	fs := S.DTypes()
	return stat.Sqrt(fs)
}
