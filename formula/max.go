package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// MAX 两个序列横向对比
func MAX(S1 stat.Series, S2 any) stat.Series {
	length := S1.Len()
	var b []num.DType
	switch sx := S2.(type) {
	case stat.Series:
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
	d := num.Maximum(S1.DTypes(), b)
	return stat.ToSeries(d...)

}
