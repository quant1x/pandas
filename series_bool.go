package pandas

import (
	"reflect"
)

type SeriesBool struct {
	NDFrame
	Data []bool
}

// NewSeriesBool creates a new series with the underlying type as bool.
func NewSeriesBool(name string, vals ...interface{}) *SeriesBool {
	series := SeriesBool{
		NDFrame: NDFrame{
			name:      name,
			nilCount:  0,
			formatter: DefaultFormatter,
		},
		Data: []bool{},
	}

	series.Data = make([]bool, 0) // Warning: filled with 0.0 (not NaN)
	size := len(series.Data)
	for idx, v := range vals {
		switch val := v.(type) {
		case bool:
			series.assign(idx, size, val)
		case []bool:
			for idx, v := range val {
				series.assign(idx, size, v)
			}
		default: // 其它容错处理
			vv := reflect.ValueOf(val)
			vk := vv.Kind()
			switch vk {
			case reflect.Invalid: // {interface} nil
				series.assign(idx, size, BoolNaN)
			case reflect.Slice: // 切片, 不定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					str := AnyToBool(tv)
					series.assign(idx, size, str)
				}
			case reflect.Array: // 数组, 定长
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					str := AnyToBool(tv)
					series.assign(idx, size, str)
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
				vv := AnyToBool(val)
				series.assign(idx, size, vv)

			}

		}
	}

	// TODO: 下面这段代码需要仔细研究, 是否存在冗余
	var lVals int
	if len(vals) > 0 {
		if ss, ok := vals[0].([]bool); ok {
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

func (self *SeriesBool) assign(idx, size int, b bool) {
	//val := AnyToString(s)
	if idx < size {
		self.Data[idx] = b
	} else {
		self.Data = append(self.Data, b)
	}
}

func (self *SeriesBool) Name() string {
	return self.name
}

func (self *SeriesBool) Rename(n string) {
	self.name = n
}

func (self *SeriesBool) Type() Type {
	return SERIES_TYPE_BOOL
}

func (self *SeriesBool) Len() int {
	return len(self.Data)
}

func (self *SeriesBool) Values() any {
	return self.Data
}

func (self *SeriesBool) Empty() Series {
	return NewSeriesBool(self.name, []bool{})
}

func (self *SeriesBool) Records() []string {
	ret := make([]string, self.Len())
	for i := 0; i < self.Len(); i++ {
		e := self.Data[i]
		ret[i] = bool2String(e)
	}
	return ret
}

func (self *SeriesBool) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen)
}

func (self *SeriesBool) Subset(start, end int, opt ...any) Series {
	var d Series
	d = NewSeriesBool(self.name, self.Data[start:end])
	return d
}

func (self *SeriesBool) Repeat(x any, repeats int) Series {
	a := AnyToFloat64(x)
	data := Repeat(a, repeats)
	var d Series
	d = NewSeriesBool(self.name, data)
	return d
}

func (self *SeriesBool) Shift(periods int) Series {
	var d Series
	d = clone(self).(Series)
	return Shift[bool](&d, periods, func() bool {
		return BoolNaN
	})
}

func (self *SeriesBool) Rolling(window int) RollingWindow {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesBool) Mean() float64 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesBool) StdDev() float64 {
	//TODO implement me
	panic("implement me")
}
