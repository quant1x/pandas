package stat

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"reflect"
)

// Series
//
//	Data structure for 1-dimensional cross-sectional and time series data
//	一维横截面和时间序列数据的数据结构
type Series interface {
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

	// Floats 强制转成[]float32
	Floats() []float32
	// DTypes 强制转[]stat.DType
	DTypes() []DType
	// Ints 强制转换成整型
	Ints() []Int
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
	Apply2(f func(idx int, v any) any, args ...bool) Series
	// Logic 逻辑处理
	Logic(f func(idx int, v any) bool) []bool
	// EWM Provide exponentially weighted (EW) calculations.
	//
	//	Exactly one of ``com``, ``span``, ``halflife``, or ``alpha`` must be
	//	provided if ``times`` is not provided. If ``times`` is provided,
	//	``halflife`` and one of ``com``, ``span`` or ``alpha`` may be provided.
	EWM(alpha EW) ExponentialMovingWindow

	// Mean calculates the average value of a series
	Mean() DType
	// StdDev calculates the standard deviation of a series
	StdDev() DType
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
	Std() DType
	// Sum 计算累和
	Sum() DType
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

// NewSeries 构建一个新的Series
func NewSeries[T BaseType](data ...T) Series {
	var S Series
	values := []T{}
	if len(data) > 0 {
		values = append(values, data...)
	}
	S = NDArray[T](values)
	return S
}

// ToSeries 转换切片为Series
func ToSeries[T BaseType](data ...T) Series {
	return NDArray[T](data)
}
