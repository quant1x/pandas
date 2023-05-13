package stat

import (
	"gitee.com/quant1x/gox/vek"
	"math"
	"reflect"
	"strings"
)

var (
	// Nil2Float64 nil指针转换float64
	Nil2Float64 = float64(0)
	// Nil2Float32 nil指针转换float32
	Nil2Float32 = float32(0)
	DTypeNaN    = DType(0)
)

var (
	// IgnoreParseExceptions 忽略解析异常
	IgnoreParseExceptions bool = true
	Avx2Enabled                = false // AVX2加速开关
)

func init() {
	Nil2Float64 = math.NaN()
	// 这个转换是对的, NaN对float32也有效
	Nil2Float32 = float32(Nil2Float64)
	DTypeNaN = DType(Nil2Float64)
}

// SetAvx2Enabled 设定AVX2加速开关
func SetAvx2Enabled(enabled bool) {
	vek.SetAcceleration(enabled)
	Avx2Enabled = enabled
}

// GetAvx2Enabled 获取avx2加速状态
func GetAvx2Enabled() bool {
	return Avx2Enabled
}

// 从指针/地址提取值
// Extract value from pointer
func extraceValueFromPointer(v any) (any, bool) {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Pointer {
		if vv.IsNil() {
			return nil, true
		}
		ve := vv.Elem()
		return ve.Interface(), true
	}
	return v, false

	//kind := reflect.ValueOf(v).Kind()
	//return nil, reflect.Pointer == kind
}

// IsEmpty Code to test if string is empty
func IsEmpty(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	} else {
		return false
	}
}
