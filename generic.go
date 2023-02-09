package pandas

import (
	"gitee.com/quant1x/pandas/algorithms/avx2"
	"gitee.com/quant1x/pandas/stat"
	gs "gonum.org/v1/gonum/stat"
	"reflect"
	"sync"
)

// GenericType Series支持的所有类型
type GenericType interface {
	~bool | ~int64 | ~float32 | ~float64 | ~string
}

// NDFrame 这里本意是想做一个父类, 实际的效果是一个抽象类
type NDFrame struct {
	lock      sync.RWMutex    // 读写锁
	formatter StringFormatter // 字符串格式化工具
	name      string          // 帧名称
	type_     Type            // values元素类型
	copy_     bool            // 是否副本
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
	if frame.Type() == SERIES_TYPE_FLOAT32 && Float32IsNaN(_vi.(float32)) {
		frame.nilCount++
	} else if frame.Type() == SERIES_TYPE_FLOAT64 && Float64IsNaN(_vi.(float64)) {
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

func (self *NDFrame) Values() any {
	return self.values
}

// NaN 输出默认的NaN
func (self *NDFrame) NaN() any {
	switch self.values.(type) {
	case []bool:
		return BoolNaN
	case []string:
		return StringNaN
	case []int64:
		return Nil2Int64
	case []float32:
		return Nil2Float32
	case []float64:
		return Nil2Float64
	default:
		panic(ErrUnsupportedType)
	}
}

func (self *NDFrame) Float() []float32 {
	return ToFloat32(self)
}

// DTypes 计算以这个函数为主
func (self *NDFrame) DTypes() []stat.DType {
	return stat.Slice2DType(self.Values())
}

// AsInt 强制转换成整型
func (self *NDFrame) AsInt() []stat.Int {
	values := self.DTypes()
	fs := stat.Fill[stat.DType](values, stat.DType(0))
	ns := stat.DType2Int(fs)
	return ns
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
	} else if self.type_ == SERIES_TYPE_INT64 {
		frame = NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      0,
			values:    []int64{},
		}
	} else if self.type_ == SERIES_TYPE_FLOAT32 {
		frame = NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      0,
			values:    []float32{},
		}
	} else if self.type_ == SERIES_TYPE_FLOAT64 {
		frame = NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      0,
			values:    []float64{},
		}
	} else {
		panic(ErrUnsupportedType)
	}
	return &frame
}

func (self *NDFrame) Records() []string {
	ret := make([]string, self.Len())
	self.Apply(func(idx int, v any) {
		ret[idx] = AnyToString(v)
	})
	return ret
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
	case []float32:
		vs := Repeat(AnyToFloat32(x), repeats)
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
	case []float32:
		return Shift[float32](&d, periods, func() float32 {
			return Nil2Float32
		})
	default: //case []float64:
		return Shift[float64](&d, periods, func() float64 {
			return Nil2Float64
		})
	}
}

func (self *NDFrame) Mean() stat.DType {
	if self.Len() < 1 {
		return NaN()
	}
	fs := make([]stat.DType, 0)
	self.Apply(func(idx int, v any) {
		f := stat.Any2DType(v)
		fs = append(fs, f)
	})
	stdDev := avx2.Mean(fs)
	return stdDev
}

func (self *NDFrame) StdDev() stat.DType {
	if self.Len() < 1 {
		return NaN()
	}
	values := make([]stat.DType, self.Len())
	self.Apply(func(idx int, v any) {
		values[idx] = stat.Any2DType(v)
	})
	stdDev := gs.StdDev(values, nil)
	return stdDev
}

func (self *NDFrame) Std() stat.DType {
	if self.Len() < 1 {
		return NaN()
	}
	values := make([]stat.DType, self.Len())
	self.Apply(func(idx int, v any) {
		values[idx] = stat.Any2DType(v)
	})
	stdDev := stat.Std(values)
	return stdDev
}

func (self *NDFrame) FillNa(v any, inplace bool) Series {
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
	case []float32:
		for idx, iv := range rows {
			if Float32IsNaN(iv) && inplace {
				rows[idx] = AnyToFloat32(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if Float64IsNaN(iv) && inplace {
				rows[idx] = AnyToFloat64(v)
			}
		}
	}
	return self
}
