package pandas

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

// Type is a convenience alias that can be used for a more type safe way of
// reason and use Series types.
type Type = reflect.Kind

// Supported Series Types
const (
	SERIES_TYPE_INVAILD = reflect.Invalid // 无效类型
	SERIES_TYPE_BOOL    = reflect.Bool    // 布尔类型
	SERIES_TYPE_INT64   = reflect.Int64   // int64
	SERIES_TYPE_FLOAT32 = reflect.Float32 // float32
	SERIES_TYPE_FLOAT64 = reflect.Float64 // float64
	SERIES_TYPE_DTYPE   = SERIES_TYPE_FLOAT64
	SERIES_TYPE_STRING  = reflect.String // string
)

type Series interface {
	// Name 取得series名称
	Name() string
	// Rename renames the series.
	Rename(name string)
	// Type returns the type of data the series holds.
	// 返回series的数据类型
	Type() Type
	// Values 获得全部数据集
	Values() any

	// NaN 输出默认的NaN
	NaN() any
	// Float 强制转成[]float32
	Float() []float32
	// DTypes 强制转[]stat.DType
	DTypes() []stat.DType
	// 强制转换成整型
	AsInt() []stat.Int

	// sort.Interface

	// Len 获得行数, 实现sort.Interface接口的获取元素数量方法
	Len() int
	// Less 实现sort.Interface接口的比较元素方法
	Less(i, j int) bool
	// Swap 实现sort.Interface接口的交换元素方法
	Swap(i, j int)

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
	// RollingV1 creates new RollingWindowV1
	// Deprecated: 使用RollingAndExpandingMixin
	RollingV1(window int) RollingWindowV1
	// Rolling 序列化版本
	Rolling(param any) RollingAndExpandingMixin
	// Mean calculates the average value of a series
	Mean() stat.DType
	// StdDev calculates the standard deviation of a series
	StdDev() stat.DType
	// FillNa Fill NA/NaN values using the specified method.
	FillNa(v any, inplace bool) Series
	// Max 找出最大值
	Max() any
	// Min 找出最小值
	Min() any
	// Select 选取一段记录
	Select(r stat.ScopeLimit) Series
	// Append 增加一批记录
	Append(values ...any)
	// Apply 接受一个回调函数
	Apply(f func(idx int, v any))
	// Logic 逻辑处理
	Logic(f func(idx int, v any) bool) []bool
	// Diff 元素的第一个离散差
	Diff(param any) (s Series)
	// Ref 引用其它周期的数据
	Ref(param any) (s Series)
	// Std 计算标准差
	Std() stat.DType
	// Sum 计算累和
	Sum() stat.DType
	// EWM Provide exponentially weighted (EW) calculations.
	//
	//    Exactly one of ``com``, ``span``, ``halflife``, or ``alpha`` must be
	//    provided if ``times`` is not provided. If ``times`` is provided,
	//    ``halflife`` and one of ``com``, ``span`` or ``alpha`` may be provided.
	EWM(alpha EW) ExponentialMovingWindow
}

// NewSeries 指定类型创建序列
func NewSeries(t Type, name string, vals any) Series {
	var series Series
	if t == SERIES_TYPE_BOOL {
		series = NewSeriesWithType(SERIES_TYPE_BOOL, name, vals)
	} else if t == SERIES_TYPE_INT64 {
		series = NewSeriesWithType(SERIES_TYPE_INT64, name, vals)
	} else if t == SERIES_TYPE_STRING {
		series = NewSeriesWithType(SERIES_TYPE_STRING, name, vals)
	} else if t == SERIES_TYPE_FLOAT64 {
		series = NewSeriesWithType(SERIES_TYPE_FLOAT64, name, vals)
	} else {
		// 默认全部强制转换成float32
		series = NewSeriesWithType(SERIES_TYPE_FLOAT32, name, vals)
	}
	return series
}

//func NewSeries_old(t Type, name string, vals ...interface{}) *Series {
//	var series Series
//	if t == SERIES_TYPE_BOOL {
//		series = NewSeriesBool(name, vals...)
//	} else if t == SERIES_TYPE_INT64 {
//		series = NewSeriesInt64(name, vals...)
//	} else if t == SERIES_TYPE_STRING {
//		series = NewSeriesString(name, vals...)
//	} else {
//		// 默认全部强制转换成float64
//		series = NewSeriesFloat64(name, vals...)
//	}
//	return &series
//}

// GenericSeries 泛型方法, 构造序列, 比其它方式对类型的统一性要求更严格
func GenericSeries[T stat.GenericType](name string, values ...T) Series {
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
		case reflect.Bool:
			_type = SERIES_TYPE_BOOL
		case reflect.Int64:
			_type = SERIES_TYPE_INT64
		case reflect.Float32:
			_type = SERIES_TYPE_FLOAT32
		case reflect.Float64:
			_type = SERIES_TYPE_FLOAT64
		case reflect.String:
			_type = SERIES_TYPE_STRING
		default:
			panic(fmt.Errorf("unknown type, %+v", v))
		}
		break
	}
	return NewSeries(_type, name, values)
}
