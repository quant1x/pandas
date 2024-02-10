package stat

import "gitee.com/quant1x/gox/exception"

// Align2Series any转换成series
func Align2Series(x any, N int) Series {
	ds := []DType{}
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
		ds = Repeat[DType](DTypeNaN, N)
	case /*nil, */ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
		ds = Repeat[DType](Any2DType(v), N)
	case []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []int, []uint, []uintptr, []float32, []float64, []bool, []string:
		vd := Slice2DType(v)
		ds = Align[DType](vd, DTypeNaN, N)
	case Series:
		vd := v.DTypes()
		ds = Align[DType](vd, DTypeNaN, N)
	default:
		panic(exception.New(1, "error window"))
	}
	return NDArray[DType](ds)
}
