package stat

import (
	"gitee.com/quant1x/pandas/exception"
	"golang.org/x/exp/slices"
	"math"
)

// Shift series切片, 使用可选的时间频率按所需的周期数移动索引
func Shift[T BaseType](S []T, periods int) []T {
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

// Shift2 series切片, 使用可选的时间频率按所需的周期数移动索引
func Shift2[T GenericType](S []T, N []DType) []T {
	var d []T
	d = slices.Clone(S)
	if len(N) == 0 {
		return d
	}
	values := d
	for i, _ := range S {
		x := N[i]
		if DTypeIsNaN(x) || int(x) > i {
			values[i] = typeDefault[T]()
			continue
		}
		values[i] = S[i-int(x)]
	}

	return d
}

// Shift3 series切片, 使用可选的时间频率按所需的周期数移动索引
//
//	param不支持负值
func Shift3[T BaseType](S []T, param any) []T {
	sLen := len(S)
	var N []DType
	switch v := param.(type) {
	case int:
		N = Repeat[DType](DType(v), sLen)
	case []DType:
		N = Align(v, DTypeNaN, sLen)
	default:
		panic(exception.New(1, "error window"))
	}
	var d []T
	d = slices.Clone(S)
	if len(N) == 0 {
		return d
	}
	values := d
	for i, _ := range S {
		x := N[i]
		if DTypeIsNaN(x) || int(x) > i {
			values[i] = typeDefault[T]()
			continue
		}
		values[i] = S[i-int(x)]
	}

	return d
}
