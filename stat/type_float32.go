package stat

import (
	"fmt"
	"gitee.com/quant1x/pandas/exception"
	"github.com/mymmsc/gox/logger"
	"github.com/viterin/vek/vek32"
	"math"
	"reflect"
	"strconv"
)

const (
	errorFloat32Base = errorTypeBase + int(reflect.Float32)*100
)

const (
	MaxFloat32                  = float32(math.MaxFloat32)             // float32最大值
	MinFloat32                  = float32(math.SmallestNonzeroFloat32) // float32最小值
	StringTrue2Float32  float32 = float32(1)                           // 字符串true转float32
	StringFalse2Float32 float32 = float32(0)                           // 字符串false转float32
)

// Float32IsNaN 判断float32是否NaN
func Float32IsNaN(f float32) bool {
	return Float64IsNaN(float64(f))
}

// 普通的处理方式, 将切片强制转换成float32
func slice_any_to_float32[T Number](s []T) []float32 {
	count := len(s)
	if count == 0 {
		return []float32{}
	}
	d := make([]float32, count)
	for idx, iv := range s {
		// 强制转换
		d[idx] = float32(iv)
	}
	return d
}

// SliceToFloat32 any输入只能是一维slice或者数组
func SliceToFloat32(v any) []float32 {
	var vs []float32
	switch values := v.(type) {
	case []int8:
		return slice_any_to_float32(values)
	case []uint8:
		return slice_any_to_float32(values)
	case []int16:
		return slice_any_to_float32(values)
	case []uint16:
		return slice_any_to_float32(values)
	case []int32: // 加速
		return vek32.FromInt32(values)
	case []uint32:
		return slice_any_to_float32(values)
	case []int64: // 加速
		return vek32.FromInt64(values)
	case []uint64:
		return slice_any_to_float32(values)
	case []int:
		return slice_any_to_float32(values)
	case []uint:
		return slice_any_to_float32(values)
	case []float32: // 克隆
		//return slices.Clone(values)
		return values
	case []float64: // 加速
		return vek32.FromFloat64(values)
	case []bool:
		count := len(values)
		if count == 0 {
			return []float32{}
		}
		// 加速
		return vek32.FromBool(values)
	case []string:
		count := len(values)
		if count == 0 {
			return []float32{}
		}
		vs = make([]float32, count)
		for idx, iv := range values {
			vs[idx] = float32(AnyToFloat64(iv))
		}
	default:
		vv := reflect.ValueOf(v)
		vk := vv.Kind()
		panic(exception.New(errorFloat32Base+0, fmt.Sprintf("Unsupported type: %s", vk.String())))
	}
	return []float32{}
}

// ParseFloat32 字符串转float32
func ParseFloat32(s string, v any) float32 {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("ParseFloat32 %+v, error=%+v\n", v, err)
		}
	}()

	if IsEmpty(s) {
		// TODO:NaN是针对64位, 这样直接转换应该有问题, 需要进一步确认
		return Nil2Float32
	}
	if StringIsTrue(s) {
		return StringTrue2Float32
	} else if StringIsFalse(s) {
		return StringFalse2Float32
	}

	f, err := strconv.ParseFloat(s, 32)
	if err == nil {
		return float32(f)
	}
	if IgnoreParseExceptions {
		return Nil2Float32
	}
	_ = v.(float32) // Intentionally panic
	return Nil2Float32
}

func AnyToFloat32(v any) float32 {
	if vv, ok := extraceValueFromPointer(v); ok {
		v = vv
	}

	f := valueToNumber(v, Nil2Float32, BoolToFloat32, ParseFloat32)
	return f
}
