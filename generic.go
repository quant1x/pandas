package pandas

import (
	"gitee.com/quant1x/pandas/algorithms/avx2"
	"gonum.org/v1/gonum/stat"
	"reflect"
	"sync"
)

// GenericType Series支持的所有类型
type GenericType interface {
	~bool | ~int64 | ~float64 | ~string
}

// NDFrame 这里本意是想做一个父类
type NDFrame struct {
	lock      sync.RWMutex    // 读写锁
	formatter StringFormatter // 字符串格式化工具
	name      string          // 帧名称
	type_     Type            // values元素类型
	nilCount  int             // nil和nan的元素有多少, 这种统计在bool和int64类型中不会大于0, 只对float64及string有效
	rows      int             // 行数
	values    any             // 只能是一个一维slice, 在所有的运算中, values强制转换成float64切片

}

//"""
//N-dimensional analogue of DataFrame. Store multi-dimensional in a
//size-mutable, labeled data structure
//
//Parameters
//----------
//data : BlockManager
//axes : list
//copy : bool, default False
//"""

func NewNDFrame[E GenericType](name string, rows ...E) *NDFrame {
	frame := NDFrame{
		formatter: DefaultFormatter,
		name:      name,
		type_:     SERIES_TYPE_INVAILD,
		nilCount:  0,
		rows:      0,
		values:    []E{},
	}
	// TODO: 不知道rows是否存在全部为空的情况, 只能先创建一个空的slice
	frame.values = make([]E, 0) // Warning: filled with 0.0 (not NaN)
	// 这个地方可以放心的强制转换, E已经做了类型约束
	size := len(frame.values.([]E))
	for idx, v := range rows {
		assign(&frame, idx, size, v)
	}

	return &frame
}

// 赋值
func assign[T GenericType](frame *NDFrame, idx, size int, v T) {
	// 检测类型
	if frame.type_ == SERIES_TYPE_INVAILD {
		_type, _ := detectTypes(v)
		if _type != SERIES_TYPE_INVAILD {
			frame.type_ = _type
		}
	}
	_vv := reflect.ValueOf(v)
	_vi := _vv.Interface()
	// float和string类型有可能是NaN, 对nil和NaN进行计数
	if frame.Type() == SERIES_TYPE_FLOAT && Float64IsNaN(_vi.(float64)) {
		frame.nilCount++
	} else if frame.Type() == SERIES_TYPE_STRING && StringIsNaN(_vi.(string)) {
		frame.nilCount++
		// 以下修正string的NaN值, 统一为"NaN"
		//_rv := reflect.ValueOf(StringNaN)
		//_vv.Set(_rv) // 这样赋值会崩溃
		// TODO:值可修改条件之一: 可被寻址
		// 通过反射修改变量值的前提条件之一: 这个值必须可以被寻址, 简单地说就是这个变量必须能被修改.
		// 第一步: 通过变量v反射(v的地址)
		_vp := reflect.ValueOf(&v)
		// 第二步: 取出v地址的元素(v的值)
		_vv := _vp.Elem()
		// 判断_vv是否能被修改
		if _vv.CanSet() {
			// 修改v的值为新值
			_vv.SetString(StringNaN)
			// 执行之后, 通过debug可以看到assign入参的v已经变成了"NaN"
		}
	}
	// 确保只添加了1个元素
	if idx < size {
		frame.values.([]T)[idx] = v
	} else {
		frame.values = append(frame.values.([]T), v)
	}
	// 行数+1
	frame.rows += 1
}

// Repeat 重复生成a
func Repeat[T GenericType](a T, n int) []T {
	dst := make([]T, n)
	for i := 0; i < n; i++ {
		dst[i] = a
	}
	return dst
}

// Repeat2 重复生成a
func Repeat2[T GenericType](dst []T, a T, n int) []T {
	for i := 0; i < n; i++ {
		dst[i] = a
	}
	return dst
}

func (self *NDFrame) Name() string {
	return self.name
}

func (self *NDFrame) Rename(n string) {
	self.name = n
}

func (self *NDFrame) Type() Type {
	return self.type_
}

func (self *NDFrame) Len() int {
	return self.rows
}

func (self *NDFrame) Values() any {
	return self.values
}

func (self *NDFrame) Empty() Series {
	var frame NDFrame
	if self.type_ == SERIES_TYPE_STRING {
		frame = NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      0,
			values:    []string{},
		}
	} else if self.type_ == SERIES_TYPE_BOOL {
		frame = NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      0,
			values:    []bool{},
		}
	} else if self.type_ == SERIES_TYPE_INT {
		frame = NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      0,
			values:    []int64{},
		}
	} else if self.type_ == SERIES_TYPE_FLOAT {
		frame = NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      0,
			values:    []float64{},
		}
	} else {
		panic("无法识别的类型")
	}
	return &frame
}

func (self *NDFrame) apply(f func(idx int, v any)) {
	vv := reflect.ValueOf(self.values)
	vk := vv.Kind()
	switch vk {
	case reflect.Invalid: // {interface} nil
		//series.assign(idx, size, Nil2Float64)
	case reflect.Slice: // 切片, 不定长
		for i := 0; i < vv.Len(); i++ {
			tv := vv.Index(i).Interface()
			f(i, tv)
		}
	case reflect.Array: // 数组, 定长
		for i := 0; i < vv.Len(); i++ {
			tv := vv.Index(i).Interface()
			f(i, tv)
		}
	default:
		// 其它类型忽略
	}
}

func (self *NDFrame) Records() []string {
	ret := make([]string, self.Len())
	self.apply(func(idx int, v any) {
		ret[idx] = AnyToString(v)
	})
	return ret
}

func (self *NDFrame) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen)
}

func (self *NDFrame) Subset(start, end int, opt ...any) Series {
	// 默认不copy
	var __optCopy bool = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _cp, ok := opt[0].(bool); ok {
			__optCopy = _cp
		}
	}
	var vs any
	var rows int
	switch values := self.values.(type) {
	case []bool:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]bool, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	case []string:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]string, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	case []int64:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]int64, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	case []float64:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]float64, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	}
	frame := NDFrame{
		formatter: self.formatter,
		name:      self.name,
		type_:     self.type_,
		nilCount:  0,
		rows:      rows,
		values:    vs,
	}
	var s Series
	s = &frame
	return s
}

func (self *NDFrame) Repeat(x any, repeats int) Series {
	switch values := self.values.(type) {
	case []bool:
		_ = values
		vs := Repeat(AnyToBool(x), repeats)
		return NewNDFrame(self.name, vs...)
	case []string:
		vs := Repeat(AnyToString(x), repeats)
		return NewNDFrame(self.name, vs...)
	case []int64:
		vs := Repeat(AnyToInt64(x), repeats)
		return NewNDFrame(self.name, vs...)
	default: //case []float64:
		vs := Repeat(AnyToFloat64(x), repeats)
		return NewNDFrame(self.name, vs...)
	}
}

func (self *NDFrame) Shift(periods int) Series {
	var d Series
	d = clone(self).(Series)
	//return Shift[float64](&d, periods, func() float64 {
	//	return Nil2Float64
	//})
	switch values := self.values.(type) {
	case []bool:
		_ = values
		return Shift[bool](&d, periods, func() bool {
			return BoolNaN
		})
	case []string:
		return Shift[string](&d, periods, func() string {
			return StringNaN
		})
	case []int64:
		return Shift[int64](&d, periods, func() int64 {
			return Nil2Int64
		})
	default: //case []float64:
		return Shift[float64](&d, periods, func() float64 {
			return Nil2Float64
		})
	}
}

func (self *NDFrame) Rolling(window int) RollingWindow {
	return RollingWindow{
		window: window,
		series: self,
	}
}

func (self *NDFrame) Mean() float64 {
	if self.Len() < 1 {
		return NaN()
	}
	fs := make([]float64, 0)
	self.apply(func(idx int, v any) {
		f := AnyToFloat64(v)
		fs = append(fs, f)
	})
	stdDev := avx2.Mean(fs)
	return stdDev
}

func (self *NDFrame) StdDev() float64 {
	if self.Len() < 1 {
		return NaN()
	}
	values := make([]float64, self.Len())
	self.apply(func(idx int, v any) {
		values[idx] = AnyToFloat64(v)
	})
	stdDev := stat.StdDev(values, nil)
	return stdDev
}

func (self *NDFrame) FillNa(v any, inplace bool) {
	values := self.Values()
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if StringIsNaN(iv) && inplace {
				rows[idx] = AnyToString(v)
			}
		}
	case []int64:
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) && inplace {
				rows[idx] = AnyToInt64(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if Float64IsNaN(iv) && inplace {
				rows[idx] = AnyToFloat64(v)
			}
		}
	}
}

func (self *NDFrame) Max() any {
	values := self.Values()
	switch rows := values.(type) {
	case []string:
		max := ""
		i := 0
		for idx, iv := range rows {
			if StringIsNaN(iv) {
				continue
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return max
		}
		return StringNaN
	case []int64:
		max := int64(0)
		//i := 0
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) {
				continue
			}
			if iv > max {
				max = iv
				//i = idx
			}
			_ = idx
		}
		return max
	case []float64:
		max := float64(0)
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(iv) {
				continue
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return max
		}
		return Nil2Float64
	}
	return Nil2Float64
}

func (self *NDFrame) Min() any {
	values := self.Values()
	switch rows := values.(type) {
	case []string:
		min := ""
		i := 0
		for idx, iv := range rows {
			if StringIsNaN(iv) {
				continue
			} else if i < 1 {
				min = iv
				i += 1
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return min
		}
		return StringNaN
	case []int64:
		min := int64(0)
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) {
				continue
			} else if i < 1 {
				min = iv
				i += 1
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		return min
	case []float64:
		min := float64(0)
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(iv) {
				continue
			} else if i < 1 {
				min = iv
				i += 1
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return min
		}
		return Nil2Float64
	}
	return Nil2Float64
}