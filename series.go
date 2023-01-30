package pandas

const (
	SERIES_TYPE_FLOAT = "float64"
)

// ValueToStringFormatter is used to convert a value
// into a string. Val can be nil or the concrete
// type stored by the series.
type ValueToStringFormatter func(val interface{}) string

type Series interface {
	// Type returns the type of data the series holds.
	// 返回类型的字符串
	Type() string
	// Shift index by desired number of periods with an optional time freq.
	// 使用可选的时间频率按所需的周期数移动索引。
	Shift(periods int) *Series
	// 获得行数
	NRows() int
	Values() any
	// Repeat elements of an array.
	Repeat(x any, repeats int) *Series
}
