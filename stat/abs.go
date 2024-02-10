package stat

import (
	"gitee.com/quant1x/gox/num"
	"gitee.com/quant1x/gox/num/num32"
)

// Abs 泛型绝对值
func Abs[T BaseType](x []T) []T {
	var d any
	var v any = x
	switch xv := v.(type) {
	case []float32:
		d = num32.Abs(xv)
	case []float64:
		d = num.Abs(xv)
	case []int:
		d = __abs_go(xv)
	case []int8:
		d = __abs_go(xv)
	case []int16:
		d = __abs_go(xv)
	case []int32:
		d = __abs_go(xv)
	case []int64:
		d = __abs_go(xv)
	case []uint, []uint8, []uint16, []uint32, []uint64, []uintptr:
		d = xv
	default:
		// 其它类型原样返回
		d = xv
	}
	return d.([]T)
}

func __abs_go[T Signed | Float](x []T) []T {
	xlen := len(x)
	d := make([]T, xlen)
	for i := 0; i < xlen; i++ {
		if x[i] < 0 {
			d[i] = -x[i]
		} else {
			d[i] = x[i]
		}
	}
	return d
}
