package pandas

import "math"

type SeriesInt64 struct {
	SeriesFrame
	Data []int64
}

func NewSeriesInt64(name string, vals ...interface{}) *SeriesInt64 {
	series := SeriesInt64{
		SeriesFrame: SeriesFrame{
			name:         name,
			nilCount:     0,
			valFormatter: DefaultValueFormatter,
		},
		Data: []int64{},
	}

	series.Data = make([]int64, 0) // Warning: filled with 0.0 (not NaN)
	size := len(series.Data)
	for idx, v := range vals {
		// Special case
		if idx == 0 {
			if fs, ok := vals[0].([]float64); ok {
				for idx, v := range fs {
					val := AnyToFloat64(v)
					if isNaN(val) {
						series.nilCount++
					}
					if idx < size {
						series.Data[idx] = int64(val)
					} else {
						series.Data = append(series.Data, int64(val))
					}
				}
				break
			}
		}

		val := AnyToFloat64(v)
		if isNaN(val) {
			series.nilCount++
		}

		if idx < size {
			series.Data[idx] = int64(val)
		} else {
			series.Data = append(series.Data, int64(val))
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
			series.Data[i] = int64(math.NaN())
		}
	}

	return &series
}

func (s SeriesInt64) Rename(n string) {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Type() string {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Shift(periods int) *Series {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Len() int {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Values() any {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Repeat(x any, repeats int) *Series {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Rolling(window int) RollingWindow {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Empty() Series {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Subset(start, end int) *Series {
	//TODO implement me
	panic("implement me")
}

func (s SeriesInt64) Mean() float64 {
	//TODO implement me
	panic("implement me")
}
