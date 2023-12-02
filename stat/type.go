package stat

import (
	"math/big"
	"reflect"
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

// NumberOfCPUBitsRelated The number of CPU bits is related
// Deprecated: 不推荐使用
type NumberOfCPUBitsRelated interface {
	~int | ~uint | ~uintptr
}

// /*nil, */ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64 , bool, string
// ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int | ~uint | ~float32 | ~float64 | ~bool | ~string
// uintptr

// BaseType 基础类型
type BaseType interface {
	Integer | Float | ~string | ~bool
}

// GenericType Series支持的所有类型
// Deprecated: 不推荐使用
type GenericType interface {
	~bool | ~int32 | ~int64 | ~int | ~float32 | ~float64 | ~string
}

// StatType 可以统计的类型
// Deprecated: 不推荐使用
type StatType interface {
	~int32 | ~int64 | ~float32 | ~float64
}

type BigFloat = big.Float // 预留将来可能扩展float

// Deprecated: 不推荐使用
type Number8 interface {
	~int8 | ~uint8
}

// Deprecated: 不推荐使用
type Number16 interface {
	~int16 | ~uint16
}

// Deprecated: 不推荐使用
type Number32 interface {
	~int32 | ~uint32 | float32
}

// Deprecated: 不推荐使用
type Number64 interface {
	~int64 | ~uint64 | float64 | int | uint
}

// Deprecated: 已弃用
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
		panic(Throw(v))
	}
	return T(0)
}

// any转number
func __anyToNumber[T Number](v any) T {
	switch val := v.(type) {
	case nil: // 这个地方判断nil值
		return typeDefault[T]()
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
		return T(BoolToInt(val))
	case string:
		vt := ParseFloat64(val, v)
		if Float64IsNaN(vt) {
			td := T(0)
			if !reflect.ValueOf(td).CanFloat() {
				return td
			}
		}
		return T(vt)
	default:
		panic(Throw(v))
	}
	return T(0)
}

// any转其它类型
// 支持3个方向: any到number, any到bool, any到string
func anyToGeneric[T BaseType](v any) T {
	var d any
	var to T
	switch any(to).(type) {
	case int8:
		d = __anyToNumber[int8](v)
	case uint8:
		d = __anyToNumber[uint8](v)
	case int16:
		d = __anyToNumber[int16](v)
	case uint16:
		d = __anyToNumber[uint16](v)
	case int32:
		d = __anyToNumber[int32](v)
	case uint32:
		d = __anyToNumber[uint32](v)
	case int64:
		d = __anyToNumber[int64](v)
	case uint64:
		d = __anyToNumber[uint64](v)
	case int:
		d = __anyToNumber[int](v)
	case uint:
		d = __anyToNumber[uint](v)
	case uintptr:
		d = __anyToNumber[uintptr](v)
	case float32:
		d = __anyToNumber[float32](v)
	case float64:
		d = __anyToNumber[float64](v)
	case bool:
		d = AnyToBool(v)
	case string:
		d = AnyToString(v)
	case []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []int, []uint, []uintptr, []float32, []float64:
		// 什么也不处理, 给个默认值
		d = to
	case []bool:
		d = to
	case []string:
		d = to
	default:
		panic(Throw(v))
	}
	return d.(T)
}

// GenericParse 泛型解析
func GenericParse[T BaseType](text string) T {
	return anyToGeneric[T](text)
}
