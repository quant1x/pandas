package pandas

import (
	"fmt"
	"math"
	"reflect"
)

// Type is a convenience alias that can be used for a more type safe way of
// reason and use Series types.
type Type = reflect.Kind

// Supported Series Types
//const (
//	SERIES_TYPE_INVAILD = "unknown" // 未知类型
//	SERIES_TYPE_BOOL    = "bool"    // 布尔类型
//	SERIES_TYPE_INT     = "int"     // int64
//	SERIES_TYPE_FLOAT   = "float"   // float64
//	SERIES_TYPE_STRING  = "string"  // string
//)

const (
	SERIES_TYPE_INVAILD = reflect.Invalid // 无效类型
	SERIES_TYPE_BOOL    = reflect.Bool    // 布尔类型
	SERIES_TYPE_INT     = reflect.Int64   // int64
	SERIES_TYPE_FLOAT   = reflect.Float64 // float64
	SERIES_TYPE_STRING  = reflect.String  // string
)

// StringFormatter is used to convert a value
// into a string. Val can be nil or the concrete
// type stored by the series.
type StringFormatter func(val interface{}) string

type Series interface {
	// Name 取得series名称
	Name() string
	// Rename renames the series.
	Rename(name string)
	// Type returns the type of data the series holds.
	// 返回series的数据类型
	Type() Type
	// Len 获得行数
	Len() int
	// Values 获得全部数据集
	Values() any
	// Empty returns an empty Series of the same type
	Empty() Series
	// Copy 复制
	Copy() Series
	// Records returns the elements of a Series as a []string
	Records() []string
	// Subset 获取子集
	Subset(start, end int, opt ...any) Series
	// Repeat elements of an array.
	Repeat(x any, repeats int) Series
	// Shift index by desired number of periods with an optional time freq.
	// 使用可选的时间频率按所需的周期数移动索引.
	Shift(periods int) Series
	// Rolling creates new RollingWindow
	Rolling(window int) RollingWindow
	// Rolling2 序列化版本
	Rolling2(param any) RollingAndExpandingMixin
	// Mean calculates the average value of a series
	Mean() float64
	// StdDev calculates the standard deviation of a series
	StdDev() float64
	// FillNa Fill NA/NaN values using the specified method.
	FillNa(v any, inplace bool)
	// Max 找出最大值
	Max() any
	// Min 找出最小值
	Min() any
	// Select 选取一段记录
	Select(r Range) Series
}

// NewSeries 指定类型创建序列
func NewSeries(t Type, name string, vals any) Series {
	var series Series
	if t == SERIES_TYPE_BOOL {
		series = NewSeriesWithType(SERIES_TYPE_BOOL, name, vals)
	} else if t == SERIES_TYPE_INT {
		series = NewSeriesWithType(SERIES_TYPE_INT, name, vals)
	} else if t == SERIES_TYPE_STRING {
		series = NewSeriesWithType(SERIES_TYPE_STRING, name, vals)
	} else {
		// 默认全部强制转换成float64
		series = NewSeriesWithType(SERIES_TYPE_FLOAT, name, vals)
	}
	return series
}

func NewSeries_old(t Type, name string, vals ...interface{}) *Series {
	var series Series
	if t == SERIES_TYPE_BOOL {
		series = NewSeriesBool(name, vals...)
	} else if t == SERIES_TYPE_INT {
		series = NewSeriesInt64(name, vals...)
	} else if t == SERIES_TYPE_STRING {
		series = NewSeriesString(name, vals...)
	} else {
		// 默认全部强制转换成float64
		series = NewSeriesFloat64(name, vals...)
	}
	return &series
}

// GenericSeries 泛型方法, 构造序列, 比其它方式对类型的统一性要求更严格
func GenericSeries[T GenericType](name string, values ...T) Series {
	// 第一遍, 确定类型, 找到第一个非nil的值
	var _type Type = SERIES_TYPE_STRING
	for _, v := range values {
		// 泛型处理这里会出现一个错误, invalid operation: v == nil (mismatched types T and untyped nil)
		//if v == nil {
		//	continue
		//}
		vv := reflect.ValueOf(v)
		vk := vv.Kind()
		switch vk {
		//case reflect.Invalid: // {interface} nil
		//	series.assign(idx, size, Nil2Float64)
		//case reflect.Slice: // 切片, 不定长
		//	for i := 0; i < vv.Len(); i++ {
		//		tv := vv.Index(i).Interface()
		//		str := AnyToFloat64(tv)
		//		series.assign(idx, size, str)
		//	}
		//case reflect.Array: // 数组, 定长
		//	for i := 0; i < vv.Len(); i++ {
		//		tv := vv.Index(i).Interface()
		//		av := AnyToFloat64(tv)
		//		series.assign(idx, size, av)
		//	}
		//case reflect.Struct: // 忽略结构体
		//	continue
		//default:
		//	vv := AnyToFloat64(val)
		//	series.assign(idx, size, vv)
		case reflect.Bool:
			_type = SERIES_TYPE_BOOL
		case reflect.Int64:
			_type = SERIES_TYPE_INT
		case reflect.Float64:
			_type = SERIES_TYPE_FLOAT
		case reflect.String:
			_type = SERIES_TYPE_STRING
		default:
			panic(fmt.Errorf("unknown type, %+v", v))
		}
		break
	}
	return NewSeries(_type, name, values)
}

func detectTypes[T GenericType](v T) (Type, any) {
	var _type = SERIES_TYPE_STRING
	vv := reflect.ValueOf(v)
	vk := vv.Kind()
	switch vk {
	case reflect.Invalid:
		_type = SERIES_TYPE_INVAILD
	case reflect.Bool:
		_type = SERIES_TYPE_BOOL
	case reflect.Int64:
		_type = SERIES_TYPE_INT
	case reflect.Float64:
		_type = SERIES_TYPE_FLOAT
	case reflect.String:
		_type = SERIES_TYPE_STRING
	default:
		panic(fmt.Errorf("unknown type, %+v", v))
	}
	return _type, vv.Interface()
}

// Shift series切片, 使用可选的时间频率按所需的周期数移动索引
func Shift[T GenericType](s *Series, periods int, cbNan func() T) Series {
	var d Series
	d = clone(*s).(Series)
	if periods == 0 {
		return d
	}

	values := d.Values().([]T)

	var (
		naVals []T
		dst    []T
		src    []T
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
		naVals[i] = cbNan()
	}
	_ = naVals
	return d
}

// FillNa 填充NaN的元素为v
// inplace为真是修改series元素的值
// 如果v和Values()返回值的slice类型不一致就会panic
func FillNa[T GenericType](s *NDFrame, v T, inplace bool) *NDFrame {
	values := s.Values()
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if StringIsNaN(iv) && inplace {
				rows[idx] = AnyToString(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if Float64IsNaN(iv) && inplace {
				rows[idx] = AnyToFloat64(v)
			}
		}
	}
	return s
}
