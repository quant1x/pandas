package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Equal 比较相等
func Equal[T BaseType](x, y []T) []bool {
	return binaryOperations2[T, bool](x, y, vek32.Eq, vek.Eq, __eq_go[T])
}

func __eq_go[T BaseType](x, y []T) []bool {
	length := len(x)
	d := make([]bool, length)
	for i := 0; i < length; i++ {
		if x[i] == y[i] {
			d[i] = true
		}
	}
	return d
}

//// CompareGt 比较 v > x
//func CompareGt(v []DType, x any) []bool {
//	return __compare(v, x, vek.Gt)
//}
//
//// CompareGte 比较 v >= x
//func CompareGte(v []DType, x any) []bool {
//	return __compare(v, x, vek.Gte)
//}
//
//// CompareLt 比较 v < x
//func CompareLt(v []DType, x any) []bool {
//	return __compare(v, x, vek.Lt)
//}
//
//// CompareLte 比较 v <= x
//func CompareLte(v []DType, x any) []bool {
//	return __compare(v, x, vek.Lte)
//}
//
//// __compare 比较 v 和 x
//func __compare(v []DType, x any, comparator func(x, y []float64) []bool) []bool {
//	vlen := len(v)
//
//	// 处理默认值
//	defaultValue := DType(0)
//	var X []DType
//	switch vx := x.(type) {
//	case int:
//		X = Repeat[DType](DType(vx), vlen)
//	case []DType:
//		xlen := len(vx)
//		if vlen < xlen {
//			vlen = xlen
//		}
//		X = Align[DType](vx, defaultValue, vlen)
//	case Series:
//		vs := vx.DTypes()
//		xlen := len(vs)
//		if vlen < xlen {
//			vlen = xlen
//		}
//		X = Align(vs, defaultValue, vlen)
//	default:
//		panic(exception.New(1, "error window"))
//	}
//	return comparator(v, X)
//}
