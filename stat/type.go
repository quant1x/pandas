package stat

import (
	"math"
	"math/big"
	"reflect"
)

// GenericType Series支持的所有类型
type GenericType interface {
	~bool | ~int32 | ~int64 | ~int | ~float32 | ~float64 | ~string
}

type Float interface {
	~float32 | ~float64
}

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

type Integer interface {
	Number8 | Number16 | Number32 | Number64
}

// Number int和uint的长度取决于CPU是多少位
type Number interface {
	Integer | Float
}

// 随便输入一个什么值
func typeDefault[T StatType](x T) T {
	xv := reflect.ValueOf(x)
	xk := xv.Kind()
	switch xk {
	case reflect.Int32, reflect.Int64:
		return T(0)
	case reflect.Float32, reflect.Float64:
		return T(math.NaN())
	}
	return T(0)
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
