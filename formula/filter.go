package formula

import (
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/exception"
)

// FILTER 过滤连续出现的信号
//
//	用法:
//	FILTER(X,N):X满足条件后,将其后N周期内的数据置为0,N为常量.
//	例如:
//	FILTER(CLOSE>OPEN,5)查找阳线,5天内再次出现的阳线不被记录在内
func FILTER(S stat.Series, N any) stat.Series {
	var W []stat.DType
	switch v := N.(type) {
	case int:
		W = stat.Repeat[stat.DType](stat.DType(v), S.Len())
	case stat.Series:
		vs := v.DTypes()
		W = stat.Align(vs, stat.DTypeNaN, S.Len())
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
	return stat.NDArray[stat.DType](x)
}
