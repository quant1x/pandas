package stat

import (
	"math/big"
)

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
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

// /*nil, */ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64 , bool, string
// ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int | ~uint | ~float32 | ~float64 | ~bool | ~string
// uintptr

// BaseType 基础类型
type BaseType interface {
	Integer | Float | ~string | ~bool
}

// GenericType Series支持的所有类型
type GenericType interface {
	~bool | ~int32 | ~int64 | ~int | ~float32 | ~float64 | ~string
}

// StatType 可以统计的类型
type StatType interface {
	~int32 | ~int64 | ~float32 | ~float64
}

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

type MoveType interface {
	StatType | ~bool | ~string
}

// Number int和uint的长度取决于CPU是多少位
type Number interface {
	Integer | Float
}

// 设定泛型默认值, 0或者NaN
func typeDefault[T BaseType]() T {
	var d any
	var t T
	var v any = t
	switch v.(type) {
	case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr:
		d = t
	case float32:
		d = Nil2Float32
	case float64:
		d = Nil2Float64
	case bool:
		d = false
	case string:
		d = StringNaN
	default:
		d = t
	}

	return d.(T)
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
	case uintptr:
		return T(val)
	case float32:
		return T(val)
	case float64:
		return T(val)
	case bool:
		return bool2t(val)
	case string:
		return string2t(val, v)
	default:
		panic(ErrUnsupportedType)
	}
	return T(0)
}
