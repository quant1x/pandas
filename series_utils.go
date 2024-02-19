package pandas

import (
	"gitee.com/quant1x/num"
)

// Align2Series any转换成series
//
//	N=-1, 即为全部
func Align2Series(x any, N int) Series {
	ds := []num.DType{}
	switch v := x.(type) {
	//case int:
	//	ds = Repeat(DType(v), N)
	//case []int:
	//	_N := Slice2DType(v)
	//	//nd := _N[len(_N) - 1]
	//	ds = Align[DType](_N, DTypeNaN, N)
	//case []DType:
	//	ds = Align(v, DTypeNaN, N)
	//case []T:
	//	vd := Slice2DType(v)
	//	ds = Align[DType](vd, DTypeNaN, N)
	case nil:
		ds = num.Repeat[num.DType](num.DTypeNaN, N)
	case /*nil, */ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
		ds = num.Repeat[num.DType](num.Any2DType(v), N)
	case []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []int, []uint, []uintptr, []float32, []float64, []bool, []string:
		vd := num.Slice2DType(v)
		if N == -1 {
			N = len(vd)
		}
		ds = num.Align[num.DType](vd, num.DTypeNaN, N)
	case Series:
		vd := v.DTypes()
		if N == -1 {
			N = len(vd)
		}
		ds = num.Align[num.DType](vd, num.DTypeNaN, N)
	default:
		panic(num.TypeError(v))
	}
	return vector[num.DType](ds)
}
