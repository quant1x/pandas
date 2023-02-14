package v1

import (
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

func (self *NDFrame) Ref(param any) (s Series) {
	var N []float32
	switch v := param.(type) {
	case int:
		N = stat.Repeat[float32](float32(v), self.Len())
	case []float32:
		N = stat.Align(v, Nil2Float32, self.Len())
	case Series:
		vs := v.Values()
		N = stat.SliceToFloat32(vs)
		N = stat.Align(N, Nil2Float32, self.Len())
	default:
		panic(exception.New(1, "error window"))
	}

	var d Series
	d = clone(self).(Series)
	//return Shift[float64](&d, periods, func() float64 {
	//	return Nil2Float64
	//})
	switch values := self.values.(type) {
	case []bool:
		_ = values
		return Shift2[bool](&d, N, func() bool {
			return stat.BoolNaN
		})
	case []string:
		return Shift2[string](&d, N, func() string {
			return stat.StringNaN
		})
	case []int64:
		return Shift2[int64](&d, N, func() int64 {
			return stat.Nil2Int64
		})
	case []float32:
		return Shift2[float32](&d, N, func() float32 {
			return Nil2Float32
		})
	default: //case []float64:
		return Shift2[float64](&d, N, func() float64 {
			return Nil2Float64
		})
	}

	return d
}
