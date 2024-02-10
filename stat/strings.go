package stat

import "fmt"

// StringFormatter is used to convert a value
// into a string. Val can be nil or the concrete
// type stored by the series.
type StringFormatter func(val any) string

// DefaultFormatter will return a string representation
// of the data in a particular row.
func DefaultFormatter(v any) string {
	if v == nil {
		return StringNaN
	}
	return fmt.Sprintf("%v", v)
}
