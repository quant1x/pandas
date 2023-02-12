package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
	"github.com/viterin/vek"
)

func AND(a, b []bool) []bool {
	return vek.And(a, b)
}

func OR(a, b []bool) []bool {
	return vek.Or(a, b)
}

// CompareGt 比较 v > x
func CompareGt(v []stat.DType, x any) []bool {
	return __compare(v, x, vek.Gt)
}

// CompareGte 比较 v >= x
func CompareGte(v []stat.DType, x any) []bool {
	return __compare(v, x, vek.Gte)
}

// CompareLt 比较 v < x
func CompareLt(v []stat.DType, x any) []bool {
	return __compare(v, x, vek.Lt)
}

// CompareLte 比较 v <= x
func CompareLte(v []stat.DType, x any) []bool {
	return __compare(v, x, vek.Lte)
}

// __compare 比较 v 和 x
func __compare(v []stat.DType, x any, comparator func(x, y []float64) []bool) []bool {
	vlen := len(v)

	// 处理默认值
	defaultValue := stat.DType(0)
	var X []stat.DType
	switch vx := x.(type) {
	case int:
		X = stat.Repeat[stat.DType](stat.DType(vx), vlen)
	case []stat.DType:
		xlen := len(vx)
		if vlen < xlen {
			vlen = xlen
		}
		X = stat.Align[stat.DType](vx, defaultValue, vlen)
	case pandas.Series:
		vs := vx.DTypes()
		xlen := len(vs)
		if vlen < xlen {
			vlen = xlen
		}
		X = stat.Align(vs, defaultValue, vlen)
	default:
		panic(exception.New(1, "error window"))
	}
	return comparator(v, X)
}
