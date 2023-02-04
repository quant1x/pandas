package stat

import (
	"math"
	"reflect"
)

type Float interface {
	~float32 | ~float64
}

type StatType interface {
	~int32 | ~int64 | ~float32 | ~float64
}

type MoveType interface {
	StatType | ~bool | ~string
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
