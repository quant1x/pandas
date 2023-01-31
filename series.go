package pandas

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

// Type is a convenience alias that can be used for a more type safe way of
// reason and use Series types.
type Type = string

type SeriesGenericType interface {
	~bool | ~int64 | ~float64 | ~string
}

// Supported Series Types
const (
	SERIES_TYPE_BOOL   = "bool"
	SERIES_TYPE_INT    = "int"
	SERIES_TYPE_FLOAT  = "float"
	SERIES_TYPE_STRING = "string"
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
	Type() Type
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
		return StringNaN
	}
	return fmt.Sprintf("%v", v)
}
