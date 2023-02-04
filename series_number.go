package pandas

import (
	"gitee.com/quant1x/pandas/algorithms/avx2"
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
	~int64 | ~uint64 | float64
}

type Float interface {
	~float32 | ~float64
}

// Number int和uint的长度取决于CPU是多少位
type Number interface {
	Number8 | Number16 | Number32 | Number64 | Float | int | uint
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
	s := avx2.Mean(d)
	return float64(s)
}

// any转number
func value_to_number[T Number](v any, nil2t T, bool2t func(b bool) T, string2t func(s string, v any) T) T {
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
func point_to_number[T Number](v any, nil2t T, bool2t func(b bool) T, string2t func(s string, v any) T) T {
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
