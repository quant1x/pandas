package pandas

import (
	"gitee.com/quant1x/pandas/algorithms"
	"gitee.com/quant1x/pandas/algorithms/avx2"
	//"github.com/huandu/go-clone"
	"gonum.org/v1/gonum/stat"
	"reflect"
)

type SeriesFloat64 struct {
	NDFrame
	Data []float64
}

func NewSeriesFloat64(name string, vals ...interface{}) *SeriesFloat64 {
	series := SeriesFloat64{
		NDFrame: NDFrame{
			name:      name,
			nilCount:  0,
			formatter: DefaultFormatter,
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

func (self *SeriesFloat64) Name() string {
	return self.name
}

func (self *SeriesFloat64) Rename(n string) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	self.name = n
}

// Type returns the type of data the series holds.
func (self *SeriesFloat64) Type() Type {
	return SERIES_TYPE_FLOAT
}

func (self *SeriesFloat64) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return len(self.Data)
}

func (self *SeriesFloat64) Shift(periods int) Series {
	var d Series
	d = clone(self).(Series)
	return Shift[float64](&d, periods, func() float64 {
		return Nil2Float
	})
}

// deprecated: 不推荐使用
func (self *SeriesFloat64) oldShift(periods int) *Series {
	var d Series
	d = clone(self).(Series)
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
		naVals[i] = NaN()
	}

	return &d
}

func (self *SeriesFloat64) Values() any {
	return self.Data
}

func (self *SeriesFloat64) Repeat(x any, repeats int) Series {
	a := AnyToFloat64(x)
	data := avx2.Repeat(a, repeats)
	//data := Repeat(a, repeats)
	var d Series
	d = NewSeriesFloat64(self.name, data)
	return d
}

// Empty returns an empty Series of the same type
func (self *SeriesFloat64) Empty() Series {
	return NewSeriesFloat64(self.name, []float64{})
}

// Records returns the elements of a Series as a []string
func (self *SeriesFloat64) Records() []string {
	ret := make([]string, self.Len())
	for i := 0; i < self.Len(); i++ {
		//e := self.elements.Elem(i)
		e := self.Data[i]
		ret[i] = float2String(e)
	}
	return ret
}

func (self *SeriesFloat64) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen)
}

func (self *SeriesFloat64) Subset(start, end int, opt ...any) Series {
	var d Series
	d = NewSeriesFloat64(self.name, self.Data[start:end])
	return d
}

// Rolling creates new RollingWindow
func (self *SeriesFloat64) Rolling(window int) RollingWindow {
	return RollingWindow{
		window: window,
		series: self,
	}
}

// Mean calculates the average value of a series
func (self *SeriesFloat64) Mean() float64 {
	if self.Len() < 1 {
		return NaN()
	}
	stdDev := avx2.Mean(self.Data)
	return stdDev
}

func (self *SeriesFloat64) StdDev() float64 {
	if self.Len() < 1 {
		return NaN()
	}
	values := self.Values().([]float64)
	stdDev := stat.StdDev(values, nil)
	return stdDev
}
