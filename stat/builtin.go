package stat

import (
	"github.com/viterin/vek"
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
)

// 初始化 avx2
// 可以参考另一个实现库 gonum.org/v1/gonum/stat
func init() {
	// 开启加速选项
	vek.SetAcceleration(true)
	Nil2Float64 = math.NaN()
	// 这个转换是对的, NaN对float32也有效
	Nil2Float32 = float32(Nil2Float64)
	DTypeNaN = DType(Nil2Float64)
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
