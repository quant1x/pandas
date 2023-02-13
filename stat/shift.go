package stat

import (
	"golang.org/x/exp/slices"
	"math"
)

// Shift series切片, 使用可选的时间频率按所需的周期数移动索引
func Shift[T GenericType](S []T, periods int) []T {
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
