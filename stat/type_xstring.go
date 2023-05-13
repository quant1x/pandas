package stat

import (
	"fmt"
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/gox/logger"
	"reflect"
	"strconv"
	"strings"
)

const (
	StringNaN    = "NaN"   // 字符串NaN
	Nil2String   = "NaN"   // nil指针转string
	True2String  = "true"  // true转string
	False2String = "false" // false转string
)

var (
	// PossibleNaOfString 有可能出现的NaN字符串的全部选项
	PossibleNaOfString = []string{"NA", "NaN", "nan", "<nil>"}
)

// StringIsNaN 判断字符串是否NaN
func StringIsNaN(s string) bool {
	s = strings.TrimSpace(s)
	if strings.ToLower(s) == "nan" {
		return true
	}
	return false
}

func __anyToString(v any) string {
	switch val := v.(type) {
	case nil:
		return Nil2String
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case int64:
		return strconv.FormatInt(int64(val), 10)
	case uint64:
		return strconv.FormatUint(uint64(val), 10)
	case int:
		return strconv.Itoa(val)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uintptr:
		return strconv.FormatUint(uint64(val), 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'G', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(val), 'G', -1, 64)
	case bool:
		return BoolToString(val)
	case string:
		return val
	default:
		logger.Errorf("%s, error=The type is not recognized\n", v)
		_ = v.(string) // Intentionally panic
		return Nil2String
	}
}

func __printString(v any) string {
	switch val := v.(type) {
	case nil:
		return Nil2String
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case int64:
		return strconv.FormatInt(int64(val), 10)
	case uint64:
		return strconv.FormatUint(uint64(val), 10)
	case int:
		return strconv.Itoa(val)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uintptr:
		return strconv.FormatUint(uint64(val), 10)
	case float32, float64:
		return fmt.Sprintf("%.3f", val)
	case bool:
		return BoolToString(val)
	case string:
		return val
	default:
		logger.Errorf("%s, error=The type is not recognized\n", v)
		_ = v.(string) // Intentionally panic
		return Nil2String
	}
}

// AnyToString any转string
func AnyToString(v any) string {
	if vv, ok := extraceValueFromPointer(v); ok {
		v = vv
	}

	s := __anyToString(v)
	return s
}

func PrintString(v any) string {
	if vv, ok := extraceValueFromPointer(v); ok {
		v = vv
	}

	s := __printString(v)
	return s
}

func slice_any_to_string[T BaseType](s []T) []string {
	count := len(s)
	if count == 0 {
		return []string{}
	}
	d := make([]string, count)
	for idx, iv := range s {
		d[idx] = anyToGeneric[string](iv)
	}
	return d
}

// SliceToString any输入只能是一维slice或者数组
func SliceToString(v any) []string {
	switch values := v.(type) {
	case []int8:
		return slice_any_to_string(values)
	case []uint8:
		return slice_any_to_string(values)
	case []int16:
		return slice_any_to_string(values)
	case []uint16:
		return slice_any_to_string(values)
	case []int32:
		return slice_any_to_string(values)
	case []uint32:
		return slice_any_to_string(values)
	case []int64:
		return slice_any_to_string(values)
	case []uint64:
		return slice_any_to_string(values)
	case []int:
		return slice_any_to_string(values)
	case []uint:
		return slice_any_to_string(values)
	case []float32:
		return slice_any_to_string(values)
	case []float64:
		return slice_any_to_string(values)
	case []bool:
		return slice_any_to_string(values)
	case []string:
		return values
	default:
		vv := reflect.ValueOf(v)
		vk := vv.Kind()
		panic(exception.New(errorFloat64Base+0, fmt.Sprintf("Unsupported type: %s", vk.String())))
	}
	return []string{}
}
