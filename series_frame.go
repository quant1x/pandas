package pandas

import (
	"github.com/huandu/go-clone"
	"math"
	"sync"
)

type SeriesFrame struct {
	valFormatter ValueToStringFormatter
	lock         sync.RWMutex
	name         string
	nilCount     int
	elements     any
}

func NewSeriesFrame(name string, vals ...interface{}) *SeriesFrame {
	series := SeriesFloat64{
		SeriesFrame: SeriesFrame{
			name:         name,
			nilCount:     0,
			valFormatter: DefaultValueFormatter,
		},
		Data: []float64{},
	}

	series.Data = make([]float64, 0) // Warning: filled with 0.0 (not NaN)
	size := len(series.Data)
	for idx, v := range vals {
		if fs, ok := v.([]float64); ok {
			for idx, v := range fs {
				val := AnyToFloat64(v)
				if isNaN(val) {
					series.nilCount++
				}
				if idx < size {
					series.Data[idx] = val
				} else {
					series.Data = append(series.Data, val)
				}
			}
			continue
		} else if fs, ok := v.([]any); ok {
			for idx, v := range fs {
				val := AnyToFloat64(v)
				if isNaN(val) {
					series.nilCount++
				}
				if idx < size {
					series.Data[idx] = val
				} else {
					series.Data = append(series.Data, val)
				}
			}
			continue
		}

		val := AnyToFloat64(v)
		if isNaN(val) {
			series.nilCount++
		}

		if idx < size {
			series.Data[idx] = val
		} else {
			series.Data = append(series.Data, val)
		}
	}

	var lVals int
	if len(vals) > 0 {
		if fs, ok := vals[0].([]float64); ok {
			lVals = len(fs)
		} else {
			lVals = len(vals)
		}
	}

	if lVals < size {
		series.nilCount = series.nilCount + size - lVals
		// Fill with NaN
		for i := lVals; i < size; i++ {
			series.Data[i] = nan()
		}
	}
	series.SeriesFrame.elements = series.Data
	return &series.SeriesFrame
}

func Shift[T ~int64 | ~float64 | ~bool | ~string](s *Series, periods int, cbNan func() T) *Series {
	var d Series
	d = clone.Clone(*s).(Series)
	if periods == 0 {
		return &d
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
	return &d
}
