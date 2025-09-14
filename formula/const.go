package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// CONST 取S最后的值为常量
func CONST(S pandas.Series) pandas.Series {
	length := S.Len()
	s := S.Float32s()
	s = num.Repeat(s[length-1], S.Len())
	return pandas.SliceToSeries(s)
}
