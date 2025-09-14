package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/num/x64"
	"gitee.com/quant1x/pandas"
)

// EQ 相等
func EQ(S1, S2 pandas.Series) []bool {
	var d []bool
	switch S1.Type() {
	case pandas.SERIES_TYPE_BOOL:
		d = num.Equal[bool](S1.Values().([]bool), S2.Values().([]bool))
	case pandas.SERIES_TYPE_STRING:
		d = num.Equal[string](S1.Values().([]string), S2.Values().([]string))
	default:
		f1 := S1.DTypes()
		f2 := S2.DTypes()
		d = num.Equal[num.DType](f1, f2)
	}
	return d
}

func EQ2(S1 []num.DType, S2 int) []bool {
	return x64.EqNumber(S1, num.DType(S2))
}

func NEQ(S1, S2 pandas.Series) []bool {
	var d []bool
	switch S1.Type() {
	case pandas.SERIES_TYPE_BOOL:
		d = num.Equal[bool](S1.Values().([]bool), S2.Values().([]bool))
	case pandas.SERIES_TYPE_STRING:
		d = num.Equal[string](S1.Values().([]string), S2.Values().([]string))
	default:
		f1 := S1.DTypes()
		f2 := S2.DTypes()
		d = num.NotEqual[num.DType](f1, f2)
	}
	return d
}

func AND[T num.Number | ~bool](a, b []T) []bool {
	return num.And(a, b)
}

func OR(a, b []bool) []bool {
	return num.Or(a, b)
}

func NOT(S pandas.Series) pandas.Series {
	x := S.Bools()
	x = num.Not(x)
	return pandas.ToSeries(x...)
}

// CompareGt 比较 v > x
func CompareGt(v []num.DType, x any) []bool {
	return __compare(v, x, x64.Gt)
}

// CompareGte 比较 v >= x
func CompareGte(v []num.DType, x any) []bool {
	return __compare(v, x, x64.Gte)
}

// CompareLt 比较 v < x
func CompareLt(v []num.DType, x any) []bool {
	return __compare(v, x, x64.Lt)
}

// CompareLte 比较 v <= x
func CompareLte(v []num.DType, x any) []bool {
	return __compare(v, x, x64.Lte)
}

// __compare 比较 v 和 x
func __compare(v []num.DType, x any, comparator func(x, y []float64) []bool) []bool {
	vlen := len(v)

	// 处理默认值
	defaultValue := num.DType(0)
	var X []num.DType
	switch vx := x.(type) {
	case int:
		X = num.Repeat[num.DType](num.DType(vx), vlen)
	case []num.DType:
		xlen := len(vx)
		if vlen < xlen {
			vlen = xlen
		}
		X = num.Align[num.DType](vx, defaultValue, vlen)
	case pandas.Series:
		vs := vx.DTypes()
		xlen := len(vs)
		if vlen < xlen {
			vlen = xlen
		}
		X = num.Align(vs, defaultValue, vlen)
	default:
		panic(num.ErrInvalidWindow)
	}
	return comparator(v, X)
}
