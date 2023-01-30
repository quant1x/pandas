package pandas

import (
	"gitee.com/quant1x/pandas/algorithms/avx2"
	"gitee.com/quant1x/pandas/algorithms/winpooh32/math"
	"github.com/huandu/go-clone"
	"github.com/viterin/vek"
	"gonum.org/v1/gonum/stat"
)

type SeriesFloat64 struct {
	SeriesFrame
	Data []float64
}

func NewSeriesFloat64(name string, vals ...interface{}) *SeriesFloat64 {
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
		// Special case
		if idx == 0 {
			if fs, ok := vals[0].([]any); ok {
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
				break
			} else if fs, ok := vals[0].([]string); ok {
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
				break
			}
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

	return &series
}

func (s *SeriesFloat64) Name() string {
	return s.name
}

func (s *SeriesFloat64) Rename(n string) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.name = n
}

// Type returns the type of data the series holds.
func (s *SeriesFloat64) Type() Type {
	return SERIES_TYPE_FLOAT
}

func (s *SeriesFloat64) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.Data)
}

func (s *SeriesFloat64) Shift(periods int) *Series {
	var d Series
	d = clone.Clone(s).(Series)
	if periods == 0 {
		return &d
	}

	values := d.Values().([]float64)

	var (
		naVals []float64
		dst    []float64
		src    []float64
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
		naVals[i] = math.NaN()
	}

	return &d
}

func (s *SeriesFloat64) Values() any {
	return s.Data
}

func (s *SeriesFloat64) Repeat(x any, repeats int) *Series {
	a := AnyToFloat64(x)

	//switch val := x.(type) {
	//case int:
	//	a = float64(val)
	//case int32:
	//	a = float64(val)
	//case int64:
	//	a = float64(val)
	//default:
	//	a = float64(val)
	//
	//}
	//
	//switch reflect.TypeOf(x).Kind() {
	//case reflect.Int, reflect.Int32, reflect.Int64:
	//	a = x.(float64)
	//case reflect.Float32, reflect.Float64:
	//	a = x.(float64)
	//}
	//if f, ok := x.(float64); ok {
	//	a = f
	//} else {
	//	a = nan()
	//}
	data := vek.Repeat_Into(s.Data, a, repeats)
	var d Series
	d = NewSeriesFloat64(s.name, data)
	return &d
}

// Empty returns an empty Series of the same type
func (s *SeriesFloat64) Empty() Series {
	return NewSeriesFloat64(s.name, []float64{})
}

// Records returns the elements of a Series as a []string
func (s *SeriesFloat64) Records() []string {
	ret := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		//e := s.elements.Elem(i)
		e := s.Data[i]
		ret[i] = float2String(e)
	}
	return ret
}

func (s *SeriesFloat64) Subset(start, end int) *Series {
	var d Series
	d = NewSeriesFloat64(s.name, s.Data[start:end])
	return &d
}

// Rolling creates new RollingWindow
func (s *SeriesFloat64) Rolling(window int) RollingWindow {
	return RollingWindow{
		window: window,
		series: s,
	}
}

// Mean calculates the average value of a series
func (s *SeriesFloat64) Mean() float64 {
	if s.Len() < 1 {
		return math.NaN()
	}
	stdDev := avx2.Mean(s.Data)
	return stdDev
}

func (s *SeriesFloat64) StdDev() float64 {
	values := s.Values().([]float64)
	stdDev := stat.StdDev(values, nil)
	return stdDev
}
