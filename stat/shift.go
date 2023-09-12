package stat

import (
	"gitee.com/quant1x/gox/exception"
	"math"
	"slices"
)

// Shift 使用可选的时间频率按所需的周期数移动索引
//
//	param 支持前后移动
func Shift[T BaseType](S []T, N any) []T {
	length := len(S)
	var _n []DType
	switch v := N.(type) {
	case Series:
		_n = v.DTypes()
	case int:
		if v == 0 {
			return S
		}
		_n = Repeat[DType](DType(v), length)
	case DType:
		if v == 0 || DTypeIsNaN(v) {
			return S
		}
		_n = Repeat[DType](DType(v), length)
	case []int:
		_n = Slice2DType(v)
		_n = Align(_n, DTypeNaN, length)
	case []DType:
		_n = Align(v, DTypeNaN, length)
	default:
		panic(exception.New(1, "error window"))
	}
	var d []T
	d = slices.Clone(S)
	values := d
	for i, _ := range S {
		x := _n[i]
		pos := int(x)
		if DTypeIsNaN(x) || i-pos >= length || i-pos < 0 {
			values[i] = typeDefault[T]()
			continue
		}
		values[i] = S[i-pos]
	}

	return d
}

// ShiftN series切片, 使用可选的时间频率按所需的周期数移动索引
// Deprecated: 不推荐使用
func ShiftN[T BaseType](S []T, periods int) []T {
	d := slices.Clone(S)
	if periods == 0 {
		return d
	}

	values := d
	var (
		naVals []T
		dst    []T
		src    []T
	)

	if shlen := int(math.Abs(float64(periods))); shlen < len(values) {
		if periods > 0 {
			naVals = values[:shlen]
			dst = values[shlen:]
			src = values
		} else {
			naVals = values[len(values)-shlen:]
			dst = values[:len(values)-shlen]
			src = values[shlen:]
		}
		copy(dst, src)
	} else {
		naVals = values
	}
	for i := range naVals {
		naVals[i] = typeDefault[T]()
	}
	_ = naVals
	return d
}
