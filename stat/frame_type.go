package stat

import (
	"reflect"
	"strings"
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

var (
	sKindInvalid       = "invalid"
	sKindBool          = "bool"
	sKindInt           = "int"
	sKindInt8          = "int8"
	sKindInt16         = "int16"
	sKindInt32         = "int32"
	sKindInt64         = "int64"
	sKindUint          = "uint"
	sKindUint8         = "uint8"
	sKindUint16        = "uint16"
	sKindUint32        = "uint32"
	sKindUint64        = "uint64"
	sKindUintptr       = "uintptr"
	sKindFloat32       = "float32"
	sKindFloat64       = "float64"
	sKindComplex64     = "complex64"
	sKindComplex128    = "complex128"
	sKindArray         = "array"
	sKindChan          = "chan"
	sKindFunc          = "func"
	sKindInterface     = "interface"
	sKindMap           = "map"
	sKindPointer       = "ptr"
	sKindSlice         = "slice"
	sKindString        = "string"
	sKindUnsafePointer = "unsafe.Pointer"
	// 缓存Kind对应关系
	mapKind = map[string]reflect.Kind{
		sKindInvalid:       reflect.Invalid,
		sKindBool:          reflect.Bool,
		sKindInt:           reflect.Int,
		sKindInt8:          reflect.Int8,
		sKindInt16:         reflect.Int16,
		sKindInt32:         reflect.Int32,
		sKindInt64:         reflect.Int64,
		sKindUint:          reflect.Uint,
		sKindUint8:         reflect.Uint8,
		sKindUint16:        reflect.Uint16,
		sKindUint32:        reflect.Uint32,
		sKindUint64:        reflect.Uint64,
		sKindUintptr:       reflect.Uintptr,
		sKindFloat32:       reflect.Float32,
		sKindFloat64:       reflect.Float64,
		sKindComplex64:     reflect.Complex64,
		sKindComplex128:    reflect.Complex128,
		sKindArray:         reflect.Array,
		sKindChan:          reflect.Chan,
		sKindFunc:          reflect.Func,
		sKindInterface:     reflect.Interface,
		sKindMap:           reflect.Map,
		sKindPointer:       reflect.Pointer,
		sKindSlice:         reflect.Slice,
		sKindString:        reflect.String,
		sKindUnsafePointer: reflect.UnsafePointer,
	}
)

// 初始化全局的私有变量
var (
	rawBool     bool    = true
	typeBool            = reflect.TypeOf([]bool{})
	rawInt32    int32   = int32(0)
	typeInt32           = reflect.TypeOf([]int32{})
	rawInt64    int64   = int64(0)
	typeInt64           = reflect.TypeOf([]int64{})
	rawFloat32  float32 = float32(0)
	typeFloat32         = reflect.TypeOf([]float32{})
	rawFloat64  float64 = float64(0)
	typeFloat64         = reflect.TypeOf([]float64{})
	typeString          = reflect.TypeOf([]string{})
)

// 从泛型检测出类型
func checkoutRawType(frame any) reflect.Kind {
	ft := reflect.TypeOf(frame)
	strType := ft.String()
	pos := strings.LastIndexByte(strType, '[')
	if pos < 0 {
		return reflect.Invalid
	}
	strType = strType[pos+1:]
	pos = strings.LastIndexByte(strType, ']')
	if pos < 0 {
		return reflect.Invalid
	}
	strType = strings.TrimSpace(strType[:pos])
	if len(strType) < 1 {
		return reflect.Invalid
	}
	if t, ok := mapKind[strType]; ok {
		return t
	}
	return reflect.Invalid
}
