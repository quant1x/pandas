package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"reflect"
)

// 这里做数组统一转换
func convert[T GenericType](s Frame, v T) {
	values := s.Values()
	rawType := checkoutRawType(values)
	values, ok := values.([]T)
	_ = rawType
	_ = ok
}

func ToFloat32(s Frame) []float32 {
	length := s.Len()
	defaultSlice := vek32.Repeat(Nil2Float32, length)
	values := s.Values()
	__type := s.Type()
	if __type == SERIES_TYPE_INVAILD {
		return defaultSlice
	} else if __type == SERIES_TYPE_BOOL {
		return vek32.FromBool(values.([]bool))
	} else if __type == SERIES_TYPE_INT64 {
		return vek32.FromInt64(values.([]int64))
	} else if __type == SERIES_TYPE_FLOAT32 {
		return values.([]float32)
	} else if __type == SERIES_TYPE_FLOAT64 {
		return vek32.FromFloat64(values.([]float64))
	} else if __type == reflect.Int32 {
		return vek32.FromInt32(values.([]int32))
	} else {
		return defaultSlice
	}
}

func ToFloat64(s Frame) []float64 {
	length := s.Len()
	defaultSlice := vek.Repeat(Nil2Float64, length)
	values := s.Values()
	__type := s.Type()
	if __type == SERIES_TYPE_INVAILD {
		return defaultSlice
	} else if __type == SERIES_TYPE_BOOL {
		return vek.FromBool(values.([]bool))
	} else if __type == SERIES_TYPE_INT64 {
		return vek.FromInt64(values.([]int64))
	} else if __type == SERIES_TYPE_FLOAT32 {
		return vek.FromFloat32(values.([]float32))
	} else if __type == SERIES_TYPE_FLOAT64 {
		return values.([]float64) // 是否复制
	} else if __type == reflect.Int32 {
		return vek.FromInt32(values.([]int32))
	} else {
		return defaultSlice
	}
}

func ToBool(s Frame) []bool {
	length := s.Len()
	defaultSlice := make([]bool, length)
	values := s.Values()
	__type := s.Type()
	if __type == SERIES_TYPE_INVAILD {
		return defaultSlice
	} else if __type == SERIES_TYPE_BOOL {
		return values.([]bool)
	} else if __type == SERIES_TYPE_INT64 {
		return __toBool(values.([]int64))
	} else if __type == SERIES_TYPE_FLOAT32 {
		return __toBool(values.([]float32))
	} else {
		return defaultSlice
	}
}

func __toBool[T Number](values []T) []bool {
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
