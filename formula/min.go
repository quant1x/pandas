package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// MIN 两个序列横向对比
func MIN(S1 pandas.Series, S2 any) pandas.Series {
	length := S1.Len()
	var b []num.DType
	switch sx := S2.(type) {
	case pandas.Series:
		b = sx.DTypes()
	case int:
		b = num.Repeat[num.DType](num.DType(sx), length)
	case num.DType:
		b = num.Repeat[num.DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	default:
		panic(num.TypeError(S2))
	}
	d := num.Minimum(S1.DTypes(), b)
	return pandas.SliceToSeries(d)

}
