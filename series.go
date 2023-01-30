package pandas

const (
	SERIES_TYPE_FLOAT = "float64"
)

// ValueToStringFormatter is used to convert a value
// into a string. Val can be nil or the concrete
// type stored by the series.
type ValueToStringFormatter func(val interface{}) string

type Series interface {
	// Rename renames the series.
	Rename(n string)
	// Type returns the type of data the series holds.
	// 返回类型的字符串
	Type() string
	// Shift index by desired number of periods with an optional time freq.
	// 使用可选的时间频率按所需的周期数移动索引。
	Shift(periods int) *Series
	// NRows 获得行数
	Len() int
	// Values 获得全部数据集
	Values() any
	// Repeat elements of an array.
	Repeat(x any, repeats int) *Series
	// Rolling creates new RollingWindow
	Rolling(window int) RollingWindow
	// Empty returns an empty Series of the same type
	Empty() Series
	Subset(start, end int) *Series
	// Mean calculates the average value of a series
	Mean() float64
}
