package stat

import "github.com/mymmsc/gox/exception"

// Align Data alignment
//
//	a 通常是默认值
func Align[T BaseType](x []T, a T, dLen int) []T {
	d := []T{}
	xLen := len(x)
	if xLen == dLen {
		return x
	} else if xLen > dLen {
		// 截断
		d = make([]T, dLen)
		copy(d, x[0:dLen])
	} else {
		// 扩展内存
		d = make([]T, dLen)
		copy(d, x)
		//avx2.RepeatAll(d[xLen:], a)
		for i := xLen; i < dLen; i++ {
			d[i] = a
		}
	}
	return d
}

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
