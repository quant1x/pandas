package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// MIN 两个序列横向对比
func MIN(S1 stat.Series, S2 any) stat.Series {
	length := S1.Len()
	var b []stat.DType
	switch sx := S2.(type) {
	case stat.Series:
		b = sx.DTypes()
	case int:
		b = stat.Repeat[stat.DType](stat.DType(sx), length)
	case stat.DType:
		b = stat.Repeat[stat.DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	default:
		panic(stat.Throw(S2))
	}
	d := stat.Minimum(S1.DTypes(), b)
	return stat.NewSeries[stat.DType](d...)

}
