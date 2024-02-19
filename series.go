package pandas

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
	"reflect"
	"strings"
)

// Type is a convenience alias that can be used for a more type safe way of
// reason and use Series types.
type Type = reflect.Kind

// Supported Series Types
const (
	SERIES_TYPE_INVAILD = reflect.Invalid     // 无效类型
	SERIES_TYPE_BOOL    = reflect.Bool        // 布尔类型
	SERIES_TYPE_INT32   = reflect.Int32       // int64
	SERIES_TYPE_INT64   = reflect.Int64       // int64
	SERIES_TYPE_FLOAT32 = reflect.Float32     // float32
	SERIES_TYPE_FLOAT64 = reflect.Float64     // float64
	SERIES_TYPE_STRING  = reflect.String      // string
	SERIES_TYPE_DTYPE   = SERIES_TYPE_FLOAT64 // link float64
)

// Series
//
//	Data structure for 1-dimensional cross-sectional and time series data
//	一维横截面和时间序列数据的数据结构
//	pandas中Series无法确定类型的情况下会使用string保存切片
type Series interface {
	String() string
	// Name 取得series名称
	Name() string
	// Rename renames the series.
	Rename(name string)
	// Type returns the type of Data the series holds.
	// 返回series的数据类型
	Type() Type
	// Values 获得全部数据集
	Values() any
	// NaN 输出默认的NaN
	NaN() any

	// DTypes 强制转[]num.DType
	DTypes() []num.DType
	// Float32s 强制转成[]float32
	Float32s() []float32
	// Float64s 强制转成[]float64
	Float64s() []float64
	// Ints 强制转换成[]int
	Ints() []int
	// Int32s 强制转换成[]int32
	Int32s() []int32
	// Int64s 强制转换成[]int64
	Int64s() []int64
	// Strings 强制转换string切片
	Strings() []string
	// Bools 强制转换成bool切片
	Bools() []bool

	// sort.Interface

	// Len 获得行数, 实现sort.Interface接口的获取元素数量方法
	Len() int
	// Less 实现sort.Interface接口的比较元素方法
	Less(i, j int) bool
	// Swap 实现sort.Interface接口的交换元素方法
	Swap(i, j int)

	// Empty returns an empty Series of the same type
	Empty(t ...Type) Series
	// Copy 复制
	Copy() Series
	// Reverse 序列反转
	Reverse() Series
	// Select 选取一段记录
	Select(r api.ScopeLimit) Series
	// Append 增加一批记录
	Append(values ...any) Series
	// Concat concatenates two series together. It will return a new Series with the
	// combined elements of both Series.
	Concat(x Series) Series

	// Records returns the elements of a Series as a []string
	Records(round ...bool) []string
	// IndexOf 取一条记录, index<0时, 从后往前取值
	IndexOf(index int, opt ...any) any
	// Subset 获取子集
	Subset(start, end int, opt ...any) Series
	// Repeat elements of an array.
	Repeat(x any, repeats int) Series
	// FillNa Fill NA/NaN values using the specified method.
	FillNa(v any, inplace bool) Series

	// Ref 引用其它周期的数据
	Ref(periods any) (s Series)
	// Shift index by desired number of periods with an optional time freq.
	//	使用可选的时间频率按所需的周期数移动索引.
	Shift(periods int) Series
	// Rolling 序列化版本
	Rolling(param any) RollingAndExpandingMixin
	// Apply 接受一个回调函数
	Apply(f func(idx int, v any))
	// Apply2 增加替换功能, 默认不替换
	Apply2(f func(idx int, v any) any, inplace ...bool) Series
	// Logic 逻辑处理
	Logic(f func(idx int, v any) bool) []bool
	// EWM Provide exponentially weighted (EW) calculations.
	//
	//	Exactly one of `com`, `span`, `halflife`, or `alpha` must be
	//	provided if `times` is not provided. If `times` is provided,
	//	`halflife` and one of `com`, `span` or `alpha` may be provided.
	EWM(alpha EW) ExponentialMovingWindow

	// Mean calculates the average value of a series
	Mean() num.DType
	// StdDev calculates the standard deviation of a series
	StdDev() num.DType
	// Max 找出最大值
	Max() any
	// ArgMax Returns the indices of the maximum values along an axis
	ArgMax() int
	// Min 找出最小值
	Min() any
	// ArgMin Returns the indices of the minimum values along an axis
	ArgMin() int
	// Diff 元素的第一个离散差
	Diff(param any) (s Series)
	// Std 计算标准差
	Std() num.DType
	// Sum 计算累和
	Sum() num.DType
	// Add 加
	Add(x any) Series
	// Sub 减
	Sub(x any) Series
	// Mul 乘
	Mul(x any) Series
	// Div 除
	Div(x any) Series
	// Eq 等于
	Eq(x any) Series
	// Neq 不等于
	Neq(x any) Series
	// Gt 大于
	Gt(x any) Series
	// Gte 大于等于
	Gte(x any) Series
	// Lt 小于
	Lt(x any) Series
	// Lte 小于等于
	Lte(x any) Series
	// And 与
	And(x any) Series
	// Or 或
	Or(x any) Series
	// Not 非
	Not() Series
}

const (
	seriesDefaultName = "x"
)

// 默认series名称
func defaultSeriesName(name ...string) string {
	if len(name) == 0 {
		return seriesDefaultName
	}
	name_ := strings.TrimSpace(name[0])
	if len(name_) == 0 {
		return seriesDefaultName
	}
	return name_
}

// DetectTypeBySlice 检测类型
func DetectTypeBySlice(arr ...any) (Type, error) {
	var hasFloat32s, hasFloat64s, hasInts, hasBools, hasStrings bool
	for _, v := range arr {
		switch value := v.(type) {
		case string:
			hasStrings = true
			continue
		case float32:
			hasFloat32s = true
			continue
		case float64:
			hasFloat64s = true
			continue
		case int, int32, int64:
			hasInts = true
			continue
		case bool:
			hasBools = true
			continue
		default:
			vv := reflect.ValueOf(v)
			vk := vv.Kind()
			switch vk {
			case reflect.Slice, reflect.Array: // 切片或数组
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					t_, err := DetectTypeBySlice(tv)
					if err == nil {
						return t_, nil
					}
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
			}
			_ = value
		}
	}

	switch {
	case hasStrings:
		return SERIES_TYPE_STRING, nil
	case hasBools:
		return SERIES_TYPE_BOOL, nil
	case hasFloat32s:
		return SERIES_TYPE_FLOAT32, nil
	case hasFloat64s:
		return SERIES_TYPE_FLOAT64, nil
	case hasInts:
		return SERIES_TYPE_INT64, nil
	default:
		return SERIES_TYPE_STRING, fmt.Errorf("couldn't detect type")
	}
}

//+----------------------------------------------------------------+
//| 切片转换的一组series函数, 数据类型确定                              |
//+----------------------------------------------------------------+

// ToSeries 转换切片为Series
func ToSeries[T num.BaseType](data ...T) Series {
	return slice2series[T](data)
}

// ToVector 转成单一切片
//
//	这种用法潜在的意图是类型明确, data可能是长度为0的切片, 但是又不想传入参数, 故而实用了默认参数的用法
func ToVector[E num.BaseType](data ...E) Series {
	return slice2series[E](data)
}

// Vector 切片转Series
func Vector[E num.BaseType](data []E) Series {
	return slice2series[E](data)
}

// SliceToSeries 切片转Series
//
//	data大概率是长度大于0的切片, 这样的函数签名是为了泛型函数不写数据类型
func SliceToSeries[E num.BaseType](data []E) Series {
	return slice2series[E](data)
}

// Convect 切片转series
//
//	存在可能的强制转换类型
func Convect[T num.BaseType, F num.BaseType](data []F) Series {
	values := num.AnyToSlice[T](data, len(data))
	return slice2series(values)
}

// 切片转Series, 这样封装的目的是在调用时不用在函数名后写类型, 由data指定类型
// eg: slice2series([]float32{1,2,3,4,5})
func slice2series[E num.BaseType](data []E) Series {
	return vector[E](data)
}

// NewSeries 模糊匹配泛型切片的匿名series, NDFrame
func NewSeries[T num.BaseType](values ...T) Series {
	return SeriesWithName[T](defaultSeriesName(), values)
}

// SeriesWithoutName 创建一个新的匿名Series
func SeriesWithoutName[E num.BaseType](values ...E) Series {
	return SeriesWithName(defaultSeriesName(), values)
}

// SeriesWithName 构建一个新的Series, NDFrame
//
//	指定类型T和名称
func SeriesWithName[T num.BaseType](name string, values []T) Series {
	frame := NDFrame{
		typ:      num.CheckoutRawType(values),
		rows:     len(values),
		nilCount: 0,
		name:     defaultSeriesName(name),
		data:     vector[T](values),
	}
	return &frame
}

//+----------------------------------------------------------------+
//| 加载数据文件时需要强制转换的一组series函数, 数据类型不确定性            |
//+----------------------------------------------------------------+

// NewSeriesWithoutType 不带类型, 创建一个新series
//
//	推导values中最适合的类型, DataFrame内部调用
func NewSeriesWithoutType(name string, values ...any) Series {
	_type, err := DetectTypeBySlice(values...)
	if err != nil {
		return nil
	}
	return NewSeriesWithType(_type, name, values...)
}

// NewSeriesWithType 指定series类型, 强制导入values
//
//	推导values中最适合的类型, DataFrame内部调用
func NewSeriesWithType(typ Type, name string, values ...any) Series {
	var vector Series
	switch typ {
	case SERIES_TYPE_BOOL:
		vector = ToSeries[bool]()
	case SERIES_TYPE_INT32:
		vector = ToSeries[int32]()
	case SERIES_TYPE_INT64:
		vector = ToSeries[int64]()
	case SERIES_TYPE_FLOAT32:
		vector = ToSeries[float32]()
	case SERIES_TYPE_FLOAT64:
		vector = ToSeries[float64]()
	default:
		vector = ToSeries[string]()
	}
	vector = vector.Append(values...)
	series := NDFrame{
		name:     name,
		typ:      typ,
		nilCount: 0,
		rows:     vector.Len(),
		data:     vector,
	}

	return &series
}
