package v1

import (
	gc "github.com/huandu/go-clone"
	"math"
	"reflect"
	"strings"
)

// 收敛统一初始化
const (
	quant1xPath = "~/.quant1x"         // quant1x默认
	tmpDir      = quant1xPath + "/tmp" // 临时路径
)

// 全局变量定义

var (
	// Nil2Float64 nil指针转换float64
	Nil2Float64 = float64(0)
	// Nil2Float32 nil指针转换float32
	Nil2Float32 = float32(0)
)

func init() {
	Nil2Float64 = math.NaN()
	// 这个转换是对的, NaN对float32也有效
	Nil2Float32 = float32(Nil2Float64)
}

// NaN returns an IEEE 754 “not-a-number” value.
func NaN() float64 {
	return math.NaN()
}

// IsNaN float64是否NaN
func IsNaN(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}

// IsEmpty Code to test if string is empty
func IsEmpty(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	} else {
		return false
	}
}

// Clone 克隆一个any
func clone(v any) any {
	return gc.Clone(v)
}

func isPoint(v any) bool {
	kind := reflect.ValueOf(v).Kind()
	return reflect.Pointer == kind
}

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
