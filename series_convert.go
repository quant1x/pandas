package pandas

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/num/x32"
	"gitee.com/quant1x/num/x64"
)

// 这里做数组统一转换
func convert[T num.GenericType](s Series, v T) {
	values := s.Values()
	rawType := num.CheckoutRawType(values)
	values, ok := values.([]T)
	_ = rawType
	_ = ok
}

// ToFloat32 转换Float32
func ToFloat32(s Series) []float32 {
	values := s.Values()
	__type := s.Type()
	switch __type {
	case SERIES_TYPE_FLOAT32:
		return values.([]float32) // TODO:是否复制
	case SERIES_TYPE_FLOAT64:
		return x32.FromFloat64(values.([]float64))
	case SERIES_TYPE_INT32:
		return x32.FromInt32(values.([]int32))
	case SERIES_TYPE_INT64:
		return x32.FromInt64(values.([]int64))
	case SERIES_TYPE_BOOL:
		return x32.FromBool(values.([]bool))
	default:
		length := s.Len()
		defaultSlice := x32.Repeat(num.Nil2Float32, length)
		return defaultSlice
	}
}

func ToFloat64(s Series) []float64 {
	values := s.Values()
	__type := s.Type()
	switch __type {
	case SERIES_TYPE_FLOAT32:
		return x64.FromFloat32(values.([]float32))
	case SERIES_TYPE_FLOAT64:
		return values.([]float64) // TODO:是否复制
	case SERIES_TYPE_INT32:
		return x64.FromInt32(values.([]int32))
	case SERIES_TYPE_INT64:
		return x64.FromInt64(values.([]int64))
	case SERIES_TYPE_BOOL:
		return x64.FromBool(values.([]bool))
	default:
		length := s.Len()
		defaultSlice := num.Repeat(num.Nil2Float64, length)
		return defaultSlice
	}
}

func ToBool(s Series) []bool {
	values := s.Values()
	__type := s.Type()
	switch __type {
	case SERIES_TYPE_FLOAT32:
		return __NumberToBool_S(values.([]float32))
	case SERIES_TYPE_FLOAT64:
		return __NumberToBool_S(values.([]float64))
	case SERIES_TYPE_INT32:
		return __NumberToBool_S(values.([]int32))
	case SERIES_TYPE_INT64:
		return __NumberToBool_S(values.([]int64))
	case SERIES_TYPE_BOOL:
		return values.([]bool)
	case SERIES_TYPE_STRING:
		return __StringToBool_S(values.([]string))
	default:
		length := s.Len()
		defaultSlice := make([]bool, length)
		return defaultSlice
	}
}

func __NumberToBool_S[T num.Number](values []T) []bool {
	length := len(values)
	vs := make([]bool, length)
	for i, v := range values {
		if v != 0 {
			vs[i] = true
		} else {
			vs[i] = false
		}
	}
	return vs
}

func __StringToBool_S(values []string) []bool {
	length := len(values)
	vs := make([]bool, length)
	for i, v := range values {
		if num.StringIsTrue(v) {
			vs[i] = true
		} else {
			vs[i] = false
		}
	}
	return vs
}

func __NumberToSeries[T num.Number](x T, n int) Series {
	s := num.Repeat[T](x, n)
	return NDArray[T](s)
}
