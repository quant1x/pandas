package pandas

import (
	"fmt"
	"gitee.com/quant1x/pandas/exception"
	"reflect"
	"strconv"
	"strings"
)

const (
	MAX_FLOAT32_PRICE = float32(9999.9999) // float32的价最大阀值触发扩展到float64
)

var (
	ErrUnsupportedType = exception.New(0, "Unsupported type")
)

func mustFloat64(f float32) bool {
	if f > MAX_FLOAT32_PRICE {
		return true
	}
	return false
}

func findTypeByString(arr []string) (Type, error) {
	var hasFloats, hasInts, hasBools, hasStrings bool
	var useFloat32, useFloat64 bool
	var stringLengthEqual = -1
	var stringLenth = -1
	for _, str := range arr {
		if str == "" || str == "NaN" {
			continue
		}
		tLen := len(str)
		if strings.HasPrefix(str, "0") {
			stringLengthEqual = 0
		}
		if stringLenth < 1 {
			if stringLengthEqual == -1 {
				stringLenth = tLen
			}
		} else if stringLengthEqual >= 0 && tLen != stringLenth {
			stringLengthEqual += 1
		}

		if _, err := strconv.Atoi(str); err == nil {
			hasInts = true
			continue
		}
		if f, err := strconv.ParseFloat(str, 64); err == nil {
			hasFloats = true
			if float32(f) < MaxFloat32 {
				if mustFloat64(float32(f)) {
					useFloat64 = true
				} else {
					useFloat32 = true
				}
			}
			continue
		}
		if str == "true" || str == "false" {
			hasBools = true
			continue
		}
		hasStrings = true
	}
	if stringLengthEqual == 0 {
		hasStrings = true
	}
	// 类型优先级, string > bool > float > int, string 为默认类型
	switch {
	case hasStrings:
		return SERIES_TYPE_STRING, nil
	case hasBools:
		return SERIES_TYPE_BOOL, nil
	case useFloat32 && !useFloat64:
		return SERIES_TYPE_FLOAT32, nil
	case hasFloats:
		return SERIES_TYPE_FLOAT64, nil
	case hasInts:
		return SERIES_TYPE_INT64, nil
	default:
		return SERIES_TYPE_STRING, fmt.Errorf("couldn't detect type")
	}

}

// 检测类型
func detectTypeBySlice(arr ...any) (Type, error) {
	var hasFloat32s, hasFloat64s, hasInts, hasBools, hasStrings bool
	for _, v := range arr {
		switch value := v.(type) {
		case string:
			hasStrings = true
			continue
		case float32:
			hasFloat32s = true
			continue
		case float64:
			hasFloat64s = true
			continue
		case int, int32, int64:
			hasInts = true
			continue
		case bool:
			hasBools = true
			continue
		default:
			vv := reflect.ValueOf(v)
			vk := vv.Kind()
			switch vk {
			case reflect.Slice, reflect.Array: // 切片或数组
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					t_, err := detectTypeBySlice(tv)
					if err == nil {
						return t_, nil
					}
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
			}
			_ = value
		}
	}

	switch {
	case hasStrings:
		return SERIES_TYPE_STRING, nil
	case hasBools:
		return SERIES_TYPE_BOOL, nil
	case hasFloat32s:
		return SERIES_TYPE_FLOAT32, nil
	case hasFloat64s:
		return SERIES_TYPE_FLOAT64, nil
	case hasInts:
		return SERIES_TYPE_INT64, nil
	default:
		return SERIES_TYPE_STRING, fmt.Errorf("couldn't detect type")
	}
}

func parseType(s string) (Type, error) {
	switch s {
	case "float", "float32":
		return SERIES_TYPE_FLOAT32, nil
	case "float64":
		return SERIES_TYPE_FLOAT64, nil
	case "int", "int64", "int32", "int16", "int8":
		return SERIES_TYPE_INT64, nil
	case "string":
		return SERIES_TYPE_STRING, nil
	case "bool":
		return SERIES_TYPE_BOOL, nil
	}
	return SERIES_TYPE_INVAILD, fmt.Errorf("type (%s) is not supported", s)
}

func detectTypes[T GenericType](v T) (Type, any) {
	var _type = SERIES_TYPE_STRING
	vv := reflect.ValueOf(v)
	vk := vv.Kind()
	switch vk {
	case reflect.Invalid:
		_type = SERIES_TYPE_INVAILD
	case reflect.Bool:
		_type = SERIES_TYPE_BOOL
	case reflect.Int64:
		_type = SERIES_TYPE_INT64
	case reflect.Float32:
		_type = SERIES_TYPE_FLOAT32
	case reflect.Float64:
		_type = SERIES_TYPE_FLOAT64
	case reflect.String:
		_type = SERIES_TYPE_STRING
	default:
		panic(fmt.Errorf("unknown type, %+v", v))
	}
	return _type, vv.Interface()
}
