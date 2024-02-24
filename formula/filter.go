package formula

import (
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// FILTER 过滤连续出现的信号
//
//	用法:
//	FILTER(X,N):X满足条件后,将其后N周期内的数据置为0,N为常量.
//	例如:
//	FILTER(CLOSE>OPEN,5)查找阳线,5天内再次出现的阳线不被记录在内
func FILTER(S pandas.Series, N any) pandas.Series {
	var W []num.DType
	switch v := N.(type) {
	case int:
		W = num.Repeat[num.DType](num.DType(v), S.Len())
	case pandas.Series:
		vs := v.DTypes()
		W = num.Align(vs, num.NaN(), S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	length := S.Len()
	x := S.DTypes()
	for i := 0; i < length; i++ {
		if x[i] != 0 {
			start := i + 1
			if start >= length {
				continue
			}
			end := i + 1 + int(W[i])
			if end >= length {
				end = length
			}
			for j := start; j < end; j++ {
				x[j] = 0
			}
		}
	}
	return pandas.SliceToSeries(x)
}
