package pandas

import (
	"reflect"
	"strings"
)

var (
	//kindNames = []string{
	//	reflect.Invalid:       "invalid",
	//	reflect.Bool:          "bool",
	//	reflect.Int:           "int",
	//	reflect.Int8:          "int8",
	//	reflect.Int16:         "int16",
	//	reflect.Int32:         "int32",
	//	reflect.Int64:         "int64",
	//	reflect.Uint:          "uint",
	//	reflect.Uint8:         "uint8",
	//	reflect.Uint16:        "uint16",
	//	reflect.Uint32:        "uint32",
	//	reflect.Uint64:        "uint64",
	//	reflect.Uintptr:       "uintptr",
	//	reflect.Float32:       "float32",
	//	reflect.Float64:       "float64",
	//	reflect.Complex64:     "complex64",
	//	reflect.Complex128:    "complex128",
	//	reflect.Array:         "array",
	//	reflect.Chan:          "chan",
	//	reflect.Func:          "func",
	//	reflect.Interface:     "interface",
	//	reflect.Map:           "map",
	//	reflect.Pointer:       "ptr",
	//	reflect.Slice:         "slice",
	//	reflect.String:        "string",
	//	reflect.UnsafePointer: "unsafe.Pointer",
	//}

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

func checkoutRawType(frame any) Type {
	ft := reflect.TypeOf(frame)
	strType := ft.String()
	pos := strings.LastIndexByte(strType, '[')
	if pos < 0 {
		return SERIES_TYPE_INVAILD
	}
	strType = strType[pos+1:]
	pos = strings.LastIndexByte(strType, ']')
	if pos < 0 {
		return SERIES_TYPE_INVAILD
	}
	strType = strings.TrimSpace(strType[:pos])
	if len(strType) < 1 {
		return SERIES_TYPE_INVAILD
	}
	if t, ok := mapKind[strType]; ok {
		return t
	}
	return SERIES_TYPE_INVAILD
}
