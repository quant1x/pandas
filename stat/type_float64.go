package stat

import (
	"fmt"
	"gitee.com/quant1x/pandas/exception"
	"github.com/mymmsc/gox/logger"
	"github.com/viterin/vek"
	"math"
	"reflect"
	"strconv"
)

const (
	errorFloat64Base = errorTypeBase + int(reflect.Float64)*100

	MaxFloat64          float64 = float64(math.MaxFloat64)             // float64最大值
	MinFloat64          float64 = float64(math.SmallestNonzeroFloat64) // float64最小值
	True2Float64        float64 = float64(1)                           // true转float64
	False2Float64       float64 = float64(0)                           // false转float64
	StringNil2Float     float64 = float64(0)                           // deprecated: 字符串空指针转float64
	StringBad2Float     float64 = float64(0)                           // deprecated: 字符串解析float64异常
	StringTrue2Float64  float64 = float64(1)                           // 字符串true转float64
	StringFalse2Float64 float64 = float64(0)                           // 字符串false转float64
)

// NaN returns an IEEE 754 “not-a-number” value.
func NaN() float64 {
	return math.NaN()
}

// Float64IsNaN 判断float64是否NaN
func Float64IsNaN(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}

func slice_any_to_float64[T Number](s []T) []float64 {
	count := len(s)
	if count == 0 {
		return []float64{}
	}
	d := make([]float64, count)
	for idx, iv := range s {
		d[idx] = float64(iv)
	}
	return d
}

// ParseFloat64 字符串转float64
// 任意组合的nan字符串都会被解析成NaN
func ParseFloat64(s string, v any) float64 {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("ParseFloat64 %+v, error=%+v\n", v, err)
		}
	}()
	if IsEmpty(s) {
		return Nil2Float64
	}
	if StringIsTrue(s) {
		return StringTrue2Float64
	} else if StringIsFalse(s) {
		return StringFalse2Float64
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f
	}
	if IgnoreParseExceptions {
		return Nil2Float64
	}
	_ = v.(float64) // Intentionally panic
	return Nil2Float64
}

func AnyToFloat64(v any) float64 {
	if vv, ok := extraceValueFromPointer(v); ok {
		v = vv
	}

	f := valueToNumber(v, Nil2Float64, BoolToFloat64, ParseFloat64)
	return f
}

// SliceToFloat64 any输入只能是一维slice或者数组
func SliceToFloat64(v any) []float64 {
	var vs []float64
	switch values := v.(type) {
	case []int8:
		return slice_any_to_float64(values)
	case []uint8:
		return slice_any_to_float64(values)
	case []int16:
		return slice_any_to_float64(values)
	case []uint16:
		return slice_any_to_float64(values)
	case []int32: // 加速
		return vek.FromInt32(values)
	case []uint32:
		return slice_any_to_float64(values)
	case []int64: // 加速
		return vek.FromInt64(values)
	case []uint64:
		return slice_any_to_float64(values)
	case []int:
		return slice_any_to_float64(values)
	case []uint:
		return slice_any_to_float64(values)
	case []float32: // 加速
		return vek.FromFloat32(values)
	case []float64: // 克隆
		//return slices.Clone(values)
		return values
	case []bool:
		count := len(values)
		if count == 0 {
			return []float64{}
		}
		// 加速
		return vek.FromBool(values)
	case []string:
		count := len(values)
		if count == 0 {
			return []float64{}
		}
		vs = make([]float64, count)
		for idx, iv := range values {
			vs[idx] = AnyToFloat64(iv)
		}
	default:
		vv := reflect.ValueOf(v)
		vk := vv.Kind()
		panic(exception.New(errorFloat64Base+0, fmt.Sprintf("Unsupported type: %s", vk.String())))
	}
	return []float64{}
}
