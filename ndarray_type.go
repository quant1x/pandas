package pandas

import (
	"reflect"
)

// Type is a convenience alias that can be used for a more type safe way of
// reason and use Series types.
type Type = reflect.Kind

// Supported Series Types
const (
	SERIES_TYPE_INVAILD = reflect.Invalid // 无效类型
	SERIES_TYPE_BOOL    = reflect.Bool    // 布尔类型
	SERIES_TYPE_INT32   = reflect.Int32   // int64
	SERIES_TYPE_INT64   = reflect.Int64   // int64
	SERIES_TYPE_FLOAT32 = reflect.Float32 // float32
	SERIES_TYPE_FLOAT64 = reflect.Float64 // float64
	SERIES_TYPE_DTYPE   = SERIES_TYPE_FLOAT64
	SERIES_TYPE_STRING  = reflect.String // string
)
