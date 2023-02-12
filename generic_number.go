package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"math/big"
)

type BigFloat = big.Float // 预留将来可能扩展float

type Number8 interface {
	~int8 | ~uint8
}

type Number16 interface {
	~int16 | ~uint16
}

type Number32 interface {
	~int32 | ~uint32 | float32
}

type Number64 interface {
	~int64 | ~uint64 | float64 | int | uint
}

// NumberOfCPUBitsRelated The number of CPU bits is related
type NumberOfCPUBitsRelated interface {
	~int | ~uint | ~uintptr
}

type Integer interface {
	Number8 | Number16 | Number32 | Number64
}

// Number int和uint的长度取决于CPU是多少位
type Number interface {
	Integer | Float
}

//type Number interface {
//	constraints.Float | constraints.Integer
//}

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
// TODO:~uintptr应该是没有应用场景
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer_old Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer_old interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
// If future releases of Go add new predeclared complex numeric types,
// this constraint will be modified to include them.
type Complex interface {
	~complex64 | ~complex128
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}

//const (
//	True2Float32        float32 = float32(1) // true转float32
//	False2Float32       float32 = float32(0) // false转float32
//	StringTrue2Float32  float32 = float32(1) // 字符串true转float32
//	StringFalse2Float32 float32 = float32(0) // 字符串false转float32
//)

// Mean gonum.org/v1/gonum/stat不支持整型, 每次都要转换有点难受啊
func Mean[T Number](x []T) float64 {
	d := numberToFloat64(x)
	s := stat.Mean(d)
	return float64(s)
}

// any转number
func valueToNumber[T Number](v any, nil2t T, bool2t func(b bool) T, string2t func(s string, v any) T) T {
	switch val := v.(type) {
	case nil: // 这个地方判断nil值
		return nil2t
	case int8:
		return T(val)
	case uint8:
		return T(val)
	case int16:
		return T(val)
	case uint16:
		return T(val)
	case int32:
		return T(val)
	case uint32:
		return T(val)
	case int64:
		return T(val)
	case uint64:
		return T(val)
	case int:
		return T(val)
	case uint:
		return T(val)
	case float32:
		return T(val)
	case float64:
		return T(val)
	case bool:
		return bool2t(val)
	case string:
		return string2t(val, v)
	}
	return T(0)
}

// 指针转number
func pointToNumber[T Number](v any, nil2t T, bool2t func(b bool) T, string2t func(s string, v any) T) T {
	switch val := v.(type) {
	case *int8:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *uint8:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *int16:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *uint16:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *int32:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *uint32:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *int64:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *uint64:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *int:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *uint:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *float32:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *float64:
		if val == nil {
			return nil2t
		}
		return T(*val)
	case *bool:
		if val == nil {
			return nil2t
		}
		return bool2t(*val)
	case *string:
		if val == nil {
			return nil2t
		}
		return string2t(*val, v)
	}
	return T(0)
}

//func anyToNumber(v any) int {
//	switch val := v.(type) {
//	case nil, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
//		// 基础类型
//		series_append(&frame, idx, size, val)
//	default:
//	}
//	return 0
//}
