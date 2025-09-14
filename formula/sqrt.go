package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// SQRT 求S的平方根
func SQRT(S pandas.Series) []num.DType {
	fs := S.DTypes()
	return num.Sqrt(fs)
}
