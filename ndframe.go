package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"reflect"
	"sync"
)

// NDFrame 这里本意是想做一个父类, 实际的效果是一个抽象类
type NDFrame struct {
	lock      sync.RWMutex         // 读写锁
	formatter stat.StringFormatter // 字符串格式化工具
	name      string               // 帧名称
	type_     stat.Type            // values元素类型
	copy_     bool                 // 是否副本
	nilCount  int                  // nil和nan的元素有多少, 这种统计在bool和int64类型中不会大于0, 只对float64及string有效
	rows      int                  // 行数
	values    any                  // 只能是一个一维slice, 在所有的运算中, values强制转换成float64切片
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

func NewNDFrame[E stat.GenericType](name string, rows ...E) *NDFrame {
	frame := NDFrame{
		formatter: stat.DefaultFormatter,
		name:      name,
		type_:     stat.SERIES_TYPE_INVAILD,
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
func ndFrameAssign[T stat.GenericType](frame *NDFrame, idx, size int, v T) {
	// 检测类型
	if frame.type_ == stat.SERIES_TYPE_INVAILD {
		_type, _ := detectTypes(v)
		if _type != stat.SERIES_TYPE_INVAILD {
			frame.type_ = _type
		}
	}
	_vv := reflect.ValueOf(v)
	_vi := _vv.Interface()
	// float和string类型有可能是NaN, 对nil和NaN进行计数
	if frame.Type() == stat.SERIES_TYPE_FLOAT32 && stat.Float32IsNaN(_vi.(float32)) {
		frame.nilCount++
	} else if frame.Type() == stat.SERIES_TYPE_FLOAT64 && stat.Float64IsNaN(_vi.(float64)) {
		frame.nilCount++
	} else if frame.Type() == stat.SERIES_TYPE_STRING && stat.StringIsNaN(_vi.(string)) {
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
			_vv.SetString(stat.StringNaN)
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

func (this *NDFrame) Type() stat.Type {
	return this.type_
}

func (this *NDFrame) Values() any {
	return this.values
}

// NaN 输出默认的NaN
func (this *NDFrame) NaN() any {
	switch this.values.(type) {
	case []bool:
		return stat.BoolNaN
	case []string:
		return stat.StringNaN
	case []int64:
		return stat.Nil2Int64
	case []float32:
		return stat.Nil2Float32
	case []float64:
		return stat.Nil2Float64
	default:
		panic(ErrUnsupportedType)
	}
}

func (this *NDFrame) Floats() []float32 {
	return stat.SliceToFloat32(this.values)
}

// DTypes 计算以这个函数为主
func (this *NDFrame) DTypes() []stat.DType {
	return stat.Slice2DType(this.Values())
}

// AsInt 强制转换成整型
func (this *NDFrame) Ints() []stat.Int {
	values := this.DTypes()
	fs := stat.Fill[stat.DType](values, stat.DType(0))
	ns := stat.DType2Int(fs)
	return ns
}

func (this *NDFrame) Strings() []string {
	return stat.SliceToString(this.Values())
}

func (this *NDFrame) Bools() []bool {
	return stat.ToBool(this)
}

func (this *NDFrame) Empty(t ...stat.Type) stat.Series {
	if len(t) > 0 {
		this.type_ = t[0]
	}
	var frame NDFrame
	if this.type_ == stat.SERIES_TYPE_STRING {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []string{},
		}
	} else if this.type_ == stat.SERIES_TYPE_BOOL {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []bool{},
		}
	} else if this.type_ == stat.SERIES_TYPE_INT64 {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []int64{},
		}
	} else if this.type_ == stat.SERIES_TYPE_FLOAT32 {
		frame = NDFrame{
			formatter: this.formatter,
			name:      this.name,
			type_:     this.type_,
			nilCount:  0,
			rows:      0,
			values:    []float32{},
		}
	} else if this.type_ == stat.SERIES_TYPE_FLOAT64 {
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
		if needRound && (t == stat.SERIES_TYPE_FLOAT32 || t == stat.SERIES_TYPE_FLOAT64) {
			ret[idx] = stat.PrintString(val)
		} else {
			ret[idx] = stat.AnyToString(val)
		}
	})
	return ret
}

func (this *NDFrame) Repeat(x any, repeats int) stat.Series {
	switch values := this.values.(type) {
	case []bool:
		_ = values
		vs := stat.Repeat(stat.AnyToBool(x), repeats)
		return NewNDFrame(this.name, vs...)
	case []string:
		vs := stat.Repeat(stat.AnyToString(x), repeats)
		return NewNDFrame(this.name, vs...)
	case []int64:
		vs := stat.Repeat(stat.AnyToInt64(x), repeats)
		return NewNDFrame(this.name, vs...)
	case []float32:
		vs := stat.Repeat(stat.AnyToFloat32(x), repeats)
		return NewNDFrame(this.name, vs...)
	default: //case []float64:
		vs := stat.Repeat(stat.AnyToFloat64(x), repeats)
		return NewNDFrame(this.name, vs...)
	}
}

func (this *NDFrame) Shift(periods int) stat.Series {
	switch values := this.values.(type) {
	case []bool:
		d := stat.Shift[bool](values, periods)
		return NewSeries(stat.SERIES_TYPE_BOOL, this.Name(), d)
	case []string:
		d := stat.Shift[string](values, periods)
		return NewSeries(stat.SERIES_TYPE_STRING, this.Name(), d)
	case []int64:
		d := stat.Shift[int64](values, periods)
		return NewSeries(stat.SERIES_TYPE_INT64, this.Name(), d)
	case []float32:
		d := stat.Shift[float32](values, periods)
		return NewSeries(stat.SERIES_TYPE_FLOAT32, this.Name(), d)
	default: //case []float64:
		d := stat.Shift[float64](values.([]float64), periods)
		return NewSeries(stat.SERIES_TYPE_FLOAT64, this.Name(), d)
	}
}

func (this *NDFrame) Mean() stat.DType {
	if this.Len() < 1 {
		return stat.NaN()
	}
	fs := make([]stat.DType, 0)
	this.Apply(func(idx int, v any) {
		f := stat.Any2DType(v)
		fs = append(fs, f)
	})
	stdDev := stat.Mean(fs)
	return stdDev
}

func (this *NDFrame) StdDev() stat.DType {
	return this.Std()
}

func (this *NDFrame) Std() stat.DType {
	if this.Len() < 1 {
		return stat.NaN()
	}
	values := make([]stat.DType, this.Len())
	this.Apply(func(idx int, v any) {
		values[idx] = stat.Any2DType(v)
	})
	stdDev := stat.Std(values)
	return stdDev
}

func (this *NDFrame) FillNa(v any, inplace bool) stat.Series {
	values := this.Values()
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if stat.StringIsNaN(iv) && inplace {
				rows[idx] = stat.AnyToString(v)
			}
		}
	case []int64:
		for idx, iv := range rows {
			if stat.Float64IsNaN(float64(iv)) && inplace {
				rows[idx] = stat.AnyToInt64(v)
			}
		}
	case []float32:
		for idx, iv := range rows {
			if stat.Float32IsNaN(iv) && inplace {
				rows[idx] = stat.AnyToFloat32(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if stat.Float64IsNaN(iv) && inplace {
				rows[idx] = stat.AnyToFloat64(v)
			}
		}
	}
	return this
}