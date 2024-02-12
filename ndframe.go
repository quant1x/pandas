package pandas

import (
	"gitee.com/quant1x/num"
	"reflect"
	"sync"
)

// NDFrame 这里本意是想做一个父类, 实际的效果是一个抽象类
type NDFrame struct {
	lock      sync.RWMutex        // 读写锁
	formatter num.StringFormatter // 字符串格式化工具
	name      string              // 帧名称
	type_     Type                // values元素类型
	copy_     bool                // 是否副本
	nilCount  int                 // nil和nan的元素有多少, 这种统计在bool和int64类型中不会大于0, 只对float64及string有效
	rows      int                 // 行数
	values    any                 // 只能是一个一维slice, 在所有的运算中, values强制转换成float64切片
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

func NewNDFrame[E num.GenericType](name string, rows ...E) *NDFrame {
	frame := NDFrame{
		formatter: num.DefaultFormatter,
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
		ndFrameAssign(&frame, idx, size, v)
	}

	return &frame
}

// 赋值
func ndFrameAssign[T num.GenericType](frame *NDFrame, idx, size int, v T) {
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
	if frame.Type() == SERIES_TYPE_FLOAT32 && num.Float32IsNaN(_vi.(float32)) {
		frame.nilCount++
	} else if frame.Type() == SERIES_TYPE_FLOAT64 && num.Float64IsNaN(_vi.(float64)) {
		frame.nilCount++
	} else if frame.Type() == SERIES_TYPE_STRING && num.StringIsNaN(_vi.(string)) {
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
			_vv.SetString(num.StringNaN)
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

func (this *NDFrame) Name() string {
	return this.name
}

func (this *NDFrame) Rename(n string) {
	this.name = n
}

func (this *NDFrame) Type() Type {
	return this.type_
}

func (this *NDFrame) Values() any {
	return this.values
}

// NaN 输出默认的NaN
func (this *NDFrame) NaN() any {
	switch this.values.(type) {
	case []bool:
		return num.BoolNaN
	case []string:
		return num.StringNaN
	case []int64:
		return num.Nil2Int64
	case []float32:
		return num.Nil2Float32
	case []float64:
		return num.Nil2Float64
	default:
		panic(ErrUnsupportedType)
	}
}

func (this *NDFrame) Floats() []float32 {
	return num.SliceToFloat32(this.values)
}

// DTypes 计算以这个函数为主
func (this *NDFrame) DTypes() []num.DType {
	return num.Slice2DType(this.Values())
}

// AsInt 强制转换成整型
func (this *NDFrame) Ints() []num.Int {
	values := this.DTypes()
	fs := num.Fill[num.DType](values, num.DType(0))
	ns := num.DType2Int(fs)
	return ns
}

func (this *NDFrame) Strings() []string {
	return num.SliceToString(this.Values())
}

func (this *NDFrame) Bools() []bool {
	return ToBool(this)
}

func (this *NDFrame) Empty(t ...Type) Series {
	if len(t) > 0 {
		this.type_ = t[0]
	}
	var frame NDFrame
	if this.type_ == SERIES_TYPE_STRING {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []string{},
		}
	} else if this.type_ == SERIES_TYPE_BOOL {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []bool{},
		}
	} else if this.type_ == SERIES_TYPE_INT64 {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []int64{},
		}
	} else if this.type_ == SERIES_TYPE_FLOAT32 {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []float32{},
		}
	} else if this.type_ == SERIES_TYPE_FLOAT64 {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []float64{},
		}
	} else {
		panic(ErrUnsupportedType)
	}
	return &frame
}

func (this *NDFrame) Records(round ...bool) []string {
	ret := make([]string, this.Len())
	needRound := false
	if len(round) > 0 {
		needRound = round[0]
	}
	t := this.Type()
	this.Apply(func(idx int, v any) {
		val := v
		if needRound && (t == SERIES_TYPE_FLOAT32 || t == SERIES_TYPE_FLOAT64) {
			ret[idx] = num.PrintString(val)
		} else {
			ret[idx] = num.AnyToString(val)
		}
	})
	return ret
}

func (this *NDFrame) Repeat(x any, repeats int) Series {
	switch values := this.values.(type) {
	case []bool:
		_ = values
		vs := num.Repeat(num.AnyToBool(x), repeats)
		return NewNDFrame(this.name, vs...)
	case []string:
		vs := num.Repeat(num.AnyToString(x), repeats)
		return NewNDFrame(this.name, vs...)
	case []int64:
		vs := num.Repeat(num.AnyToInt64(x), repeats)
		return NewNDFrame(this.name, vs...)
	case []float32:
		vs := num.Repeat(num.AnyToFloat32(x), repeats)
		return NewNDFrame(this.name, vs...)
	default: //case []float64:
		vs := num.Repeat(num.AnyToFloat64(x), repeats)
		return NewNDFrame(this.name, vs...)
	}
}

func (this *NDFrame) Shift(periods int) Series {
	switch values := this.values.(type) {
	case []bool:
		d := num.Shift[bool](values, periods)
		return NewSeries(SERIES_TYPE_BOOL, this.Name(), d)
	case []string:
		d := num.Shift[string](values, periods)
		return NewSeries(SERIES_TYPE_STRING, this.Name(), d)
	case []int64:
		d := num.Shift[int64](values, periods)
		return NewSeries(SERIES_TYPE_INT64, this.Name(), d)
	case []float32:
		d := num.Shift[float32](values, periods)
		return NewSeries(SERIES_TYPE_FLOAT32, this.Name(), d)
	default: //case []float64:
		d := num.Shift[float64](values.([]float64), periods)
		return NewSeries(SERIES_TYPE_FLOAT64, this.Name(), d)
	}
}

func (this *NDFrame) Mean() num.DType {
	if this.Len() < 1 {
		return num.NaN()
	}
	fs := make([]num.DType, 0)
	this.Apply(func(idx int, v any) {
		f := num.Any2DType(v)
		fs = append(fs, f)
	})
	stdDev := num.Mean(fs)
	return stdDev
}

func (this *NDFrame) StdDev() num.DType {
	return this.Std()
}

func (this *NDFrame) Std() num.DType {
	if this.Len() < 1 {
		return num.NaN()
	}
	values := make([]num.DType, this.Len())
	this.Apply(func(idx int, v any) {
		values[idx] = num.Any2DType(v)
	})
	stdDev := num.Std(values)
	return stdDev
}

func (this *NDFrame) FillNa(v any, inplace bool) Series {
	values := this.Values()
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if num.StringIsNaN(iv) && inplace {
				rows[idx] = num.AnyToString(v)
			}
		}
	case []int64:
		for idx, iv := range rows {
			if num.Float64IsNaN(float64(iv)) && inplace {
				rows[idx] = num.AnyToInt64(v)
			}
		}
	case []float32:
		for idx, iv := range rows {
			if num.Float32IsNaN(iv) && inplace {
				rows[idx] = num.AnyToFloat32(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if num.Float64IsNaN(iv) && inplace {
				rows[idx] = num.AnyToFloat64(v)
			}
		}
	}
	return this
}
