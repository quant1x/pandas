// Copyright 2018-20 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package pandas

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strconv"
)

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
