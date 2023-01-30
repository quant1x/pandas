package pandas

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strconv"
)

const (
	SERIES_TYPE_FLOAT = "float64"
)

// ValueToStringFormatter is used to convert a value
// into a string. Val can be nil or the concrete
// type stored by the series.
type ValueToStringFormatter func(val interface{}) string

type Series interface {
	// Name 取得series名称
	Name() string
	// Rename renames the series.
	Rename(n string)
	// Type returns the type of data the series holds.
	// 返回类型的字符串
	Type() string
	// NRows 获得行数
	Len() int
	// Values 获得全部数据集
	Values() any
	// Empty returns an empty Series of the same type
	Empty() Series
	// Records returns the elements of a Series as a []string
	Records() []string
	// Subset 获取子集
	Subset(start, end int) *Series
	// Repeat elements of an array.
	Repeat(x any, repeats int) *Series
	// Shift index by desired number of periods with an optional time freq.
	// 使用可选的时间频率按所需的周期数移动索引。
	Shift(periods int) *Series
	// Rolling creates new RollingWindow
	Rolling(window int) RollingWindow
	// Mean calculates the average value of a series
	Mean() float64
	// StdDev calculates the standard deviation of a series
	StdDev() float64
}

// DefaultIsEqualFunc is the default comparitor to determine if
// two values in the series are the same.
func DefaultIsEqualFunc(a, b interface{}) bool {
	return cmp.Equal(a, b)
}

// DefaultValueFormatter will return a string representation
// of the data in a particular row.
func DefaultValueFormatter(v interface{}) string {
	if v == nil {
		return "NaN"
	}
	return fmt.Sprintf("%v", v)
}

func AnyToFloat64(v any) float64 {
	switch val := v.(type) {
	case nil:
		return nan()
	case bool:
		if val == true {
			return float64(1)
		}
		return float64(0)
	case int:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case float64:
		return val
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			_ = v.(float64) // Intentionally panic
		}
		return f
	default:
		f, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
		if err != nil {
			_ = v.(float64) // Intentionally panic
		}
		return f
	}
}

func AnyToInt64(v any) int64 {
	f := AnyToFloat64(v)
	return int64(f)
}

func AnyToInt32(v any) int32 {
	f := AnyToFloat64(v)
	return int32(f)
}

func AnyToInt(v any) int {
	f := AnyToFloat64(v)
	return int(f)
}

func float2String(v float64) string {
	if isNaN(v) {
		return "NaN"
	}
	return fmt.Sprintf("%f", v)
}

func int2String(v int64) string {
	if isNaN(float64(v)) {
		return "NaN"
	}
	return fmt.Sprint(v)
}
