package pandas

import (
	"reflect"
)

// SeriesString 字符串类型序列
type SeriesString struct {
	NDFrame
	Data []string
}

// NewSeriesString creates a new series with the underlying type as string.
func NewSeriesString(name string, vals ...interface{}) *SeriesString {
	series := SeriesString{
		NDFrame: NDFrame{
			name:      name,
			nilCount:  0,
			formatter: DefaultFormatter,
		},
		Data: []string{},
	}

	series.Data = make([]string, 0) // Warning: filled with 0.0 (not NaN)
	size := len(series.Data)
	for idx, v := range vals {
		switch val := v.(type) {
		case string:
			series.assign(idx, size, val)
		case []string:
			for idx, v := range val {
				series.assign(idx, size, v)
			}
		default: // 其它容错处理
			vv := reflect.ValueOf(val)
			vk := vv.Kind()
			switch vk {
			case reflect.Invalid: // {interface} nil
				series.assign(idx, size, StringNaN)
			case reflect.Slice: // 切片, 不定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					str := AnyToString(tv)
					series.assign(idx, size, str)
				}
			case reflect.Array: // 数组, 定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					str := AnyToString(tv)
					series.assign(idx, size, str)
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
				vv := AnyToString(val)
				series.assign(idx, size, vv)
			}
		}
	}

	// TODO: 下面这段代码需要仔细研究, 是否存在冗余
	var lVals int
	if len(vals) > 0 {
		if ss, ok := vals[0].([]string); ok {
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

func (self *SeriesString) assign(idx, size int, s string) {
	//val := AnyToString(s)
	if StringIsNaN(s) {
		s = StringNaN
		self.nilCount++
	}
	if idx < size {
		self.Data[idx] = string(s)
	} else {
		self.Data = append(self.Data, string(s))
	}
}

func (self *SeriesString) Name() string {
	return self.name
}

func (self *SeriesString) Rename(n string) {
	self.name = n
}

func (self *SeriesString) Type() Type {
	return SERIES_TYPE_STRING
}

func (self *SeriesString) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return len(self.Data)
}

func (self *SeriesString) Values() any {
	return self.Data
}

func (self *SeriesString) Empty() Series {
	return NewSeriesString(self.name, []string{})
}

func (self *SeriesString) Records() []string {
	ret := make([]string, self.Len())
	for i := 0; i < self.Len(); i++ {
		e := self.Data[i]
		ret[i] = e
	}
	return ret
}

func (self *SeriesString) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen)
}

func (self *SeriesString) Subset(start, end int, opt ...any) Series {
	var d Series
	d = NewSeriesString(self.name, self.Data[start:end])
	return d
}

func (self *SeriesString) Repeat(x any, repeats int) Series {
	a := AnyToFloat64(x)
	data := Repeat(a, repeats)
	var d Series
	d = NewSeriesString(self.name, data)
	return d
}

func (self *SeriesString) Shift(periods int) Series {
	var d Series
	d = clone(self).(Series)
	return Shift[string](&d, periods, func() string {
		return Nil2String
	})
}

func (self *SeriesString) RollingV1(window int) RollingWindowV1 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesString) Mean() float64 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesString) StdDev() float64 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesString) FillNa(v any, inplace bool) Series {
	values := self.Values()
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if StringIsNaN(iv) && inplace {
				rows[idx] = AnyToString(v)
			}
		}
	}

	return self
}
