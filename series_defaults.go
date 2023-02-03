package pandas

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

// DefaultIsEqualFunc is the default comparitor to determine if
// two values in the series are the same.
func DefaultIsEqualFunc(a, b interface{}) bool {
	return cmp.Equal(a, b)
}

// DefaultFormatter will return a string representation
// of the data in a particular row.
func DefaultFormatter(v interface{}) string {
	if v == nil {
		return StringNaN
	}
	return fmt.Sprintf("%v", v)
}
