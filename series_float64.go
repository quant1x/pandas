package pandas

import (
	"gitee.com/quant1x/pandas/algorithms"
	"gitee.com/quant1x/pandas/algorithms/avx2"
	"github.com/huandu/go-clone"
	"gonum.org/v1/gonum/stat"
	"reflect"
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
		switch val := v.(type) {
		case float64:
			series.assign(idx, size, val)
		case []float64:
			for idx, v := range val {
				series.assign(idx, size, v)
			}
		default: // 其它容错处理
			vv := reflect.ValueOf(val)
			vk := vv.Kind()
			switch vk {
			case reflect.Invalid: // {interface} nil
				series.assign(idx, size, Nil2Float)
			case reflect.Slice: // 切片, 不定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					str := AnyToFloat64(tv)
					series.assign(idx, size, str)
				}
			case reflect.Array: // 数组, 定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					av := AnyToFloat64(tv)
					series.assign(idx, size, av)
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
				vv := AnyToFloat64(val)
				series.assign(idx, size, vv)
			}
		}
	}

	// TODO: 下面这段代码需要仔细研究, 是否存在冗余
	var lVals int
	if len(vals) > 0 {
		if ss, ok := vals[0].([]float64); ok {
			lVals = len(ss)
		} else {
			lVals = len(vals)
		}
	}

	if lVals < size {
		series.nilCount = series.nilCount + size - lVals
	}

	return &series
}

func (self *SeriesFloat64) assign(idx, size int, f float64) {
	if IsNaN(f) {
		self.nilCount++
	}
	if idx < size {
		self.Data[idx] = float64(f)
	} else {
		self.Data = append(self.Data, float64(f))
	}
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
	return Shift[float64](&d, periods, func() float64 {
		return Nil2Float
	})
}

// deprecated: 不推荐使用
func (s *SeriesFloat64) oldShift(periods int) *Series {
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

	if shlen := int(algorithms.Abs(float64(periods))); shlen < len(values) {
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
		naVals[i] = algorithms.NaN()
	}

	return &d
}

func (s *SeriesFloat64) Values() any {
	return s.Data
}

func (s *SeriesFloat64) Repeat(x any, repeats int) *Series {
	a := AnyToFloat64(x)
	//data := avx2.Repeat(a, repeats)
	data := Repeat(a, repeats)
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
		return algorithms.NaN()
	}
	stdDev := avx2.Mean(s.Data)
	return stdDev
}

func (s *SeriesFloat64) StdDev() float64 {
	values := s.Values().([]float64)
	stdDev := stat.StdDev(values, nil)
	return stdDev
}
