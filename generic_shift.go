package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"math"
)

// Shift series切片, 使用可选的时间频率按所需的周期数移动索引
func Shift[T stat.GenericType](s *Series, periods int, cbNan func() T) Series {
	var d Series
	d = clone(*s).(Series)
	if periods == 0 {
		return d
	}

	values := d.Values().([]T)

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
		naVals[i] = cbNan()
	}
	_ = naVals
	return d
}

// Shift2 series切片, 使用可选的时间频率按所需的周期数移动索引
func Shift2[T stat.GenericType](s *Series, N []float32, cbNan func() T) Series {
	var d Series
	d = clone(*s).(Series)
	if len(N) == 0 {
		return d
	}
	S := (*s).Values().([]T)
	values := d.Values().([]T)
	for i, _ := range S {
		x := N[i]
		if stat.Float32IsNaN(x) || int(x) > i {
			values[i] = cbNan()
			continue
		}
		values[i] = S[i-int(x)]
	}

	return d
}
