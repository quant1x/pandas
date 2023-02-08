package pandas

import (
	"gonum.org/v1/gonum/stat"
	"reflect"
)

type SeriesInt64 struct {
	NDFrame
	Data []int64
}

// NewSeriesInt64 creates a new series with the underlying type as int64.
func NewSeriesInt64(name string, vals ...interface{}) *SeriesInt64 {
	series := SeriesInt64{
		NDFrame: NDFrame{
			name:      name,
			nilCount:  0,
			formatter: DefaultFormatter,
		},
		Data: []int64{},
	}

	series.Data = make([]int64, 0) // Warning: filled with 0.0 (not NaN)
	size := len(series.Data)
	for idx, v := range vals {
		switch val := v.(type) {
		case int64:
			series.assign(idx, size, val)
		case []int64:
			for idx, v := range val {
				series.assign(idx, size, v)
			}
		default: // 其它容错处理
			vv := reflect.ValueOf(val)
			vk := vv.Kind()
			switch vk {
			case reflect.Invalid: // {interface} nil
				series.assign(idx, size, Int64NaN)
			case reflect.Slice: // 切片, 不定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					str := AnyToInt64(tv)
					series.assign(idx, size, str)
				}
			case reflect.Array: // 数组, 定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					str := AnyToInt64(tv)
					series.assign(idx, size, str)
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
				vv := AnyToInt64(val)
				series.assign(idx, size, vv)
			}
		}
	}

	// TODO: 下面这段代码需要仔细研究, 是否存在冗余
	var lVals int
	if len(vals) > 0 {
		if ss, ok := vals[0].([]int64); ok {
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

func (self *SeriesInt64) assign(idx, size int, n int64) {
	//if StringIsNaN(self) {
	//	self = StringNaN
	//	self.nilCount++
	//}
	if idx < size {
		self.Data[idx] = int64(n)
	} else {
		self.Data = append(self.Data, int64(n))
	}
}

func (self *SeriesInt64) Name() string {
	return self.name
}

func (self *SeriesInt64) Rename(n string) {
	self.name = n
}

func (self *SeriesInt64) Type() Type {
	return SERIES_TYPE_INT64
}

func (self *SeriesInt64) Shift(periods int) Series {
	var d Series
	d = clone(self).(Series)
	return Shift[int64](&d, periods, func() int64 {
		return Int64NaN
	})
}

func (self *SeriesInt64) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return len(self.Data)
}

func (self *SeriesInt64) Values() any {
	return self.Data
}

func (self *SeriesInt64) Repeat(x any, repeats int) Series {
	a := AnyToFloat64(x)
	data := Repeat(a, repeats)
	var d Series
	d = NewSeriesInt64(self.name, data)
	return d
}

func (self *SeriesInt64) Empty() Series {
	return NewSeriesInt64(self.name, []int64{})
}

// Records returns the elements of a Series as a []string
func (self *SeriesInt64) Records() []string {
	ret := make([]string, self.Len())
	for i := 0; i < self.Len(); i++ {
		e := self.Data[i]
		ret[i] = int64ToString(e)
	}
	return ret
}

func (self *SeriesInt64) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen)
}

func (self *SeriesInt64) Subset(start, end int, opt ...any) Series {
	var d Series
	d = NewSeriesInt64(self.name, self.Data[start:end])
	return d
}

func (self *SeriesInt64) RollingV1(window int) RollingWindowV1 {
	return RollingWindowV1{
		window: window,
		series: self,
	}
}

func (self *SeriesInt64) Mean() float64 {
	if self.Len() < 1 {
		return NaN()
	}
	stdDev := Mean(self.Data)
	return stdDev
}

func (self *SeriesInt64) StdDev() float64 {
	if self.Len() < 1 {
		return NaN()
	}
	// TODO: 每次都需要转换一次, 有没有什么好办法优化?
	d := make([]float64, self.Len())
	for i := 0; i < len(d); i++ {
		d[i] = float64(self.Data[i])
	}
	//values := self.Values().([]int64)
	stdDev := stat.StdDev(d, nil)
	return stdDev
}

// FillNa int64没有NaN
func (self *SeriesInt64) FillNa(v any, inplace bool) Series {
	values := self.Values()
	switch rows := values.(type) {
	case []int64:
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) && inplace {
				rows[idx] = AnyToInt64(v)
			}
		}
	}
	return self
}
