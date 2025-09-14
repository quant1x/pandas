package pandas

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/quant1x/num"
)

const (
	//MAX_FLOAT32_PRICE = float32(9999.9999) // float32的价最大阀值触发扩展到float64
	MAX_FLOAT32_PRICE = float32(0) // float32的价最大阀值触发扩展到float64
)

var (
	ErrUnsupportedType    = errors.New("unsupported type")
	ErrCouldNotDetectType = errors.New("couldn't detect type")
)

func mustFloat64(f float32) bool {
	if f > MAX_FLOAT32_PRICE {
		return true
	}
	return false
}

func findTypeByString(arr []string) (Type, error) {
	var hasFloats, hasInts, hasBools, hasStrings bool
	var useInt32, useInt64 bool
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
			if stringLengthEqual <= 0 {
				stringLenth = tLen
			}
		} else if stringLengthEqual >= 0 && tLen != stringLenth {
			stringLengthEqual += 1
		}
		// 整型
		if d, err := strconv.ParseInt(str, 10, 64); err == nil {
			hasInts = true
			//if int32(d) <= num.MaxInt32 {
			//	useInt32 = true
			//} else {
			//	useInt64 = true
			//}
			_ = d
		}

		// 浮点
		if f, err := strconv.ParseFloat(str, 64); err == nil {
			hasFloats = true
			if float32(f) < num.MaxFloat32 {
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
	if !hasFloats && stringLengthEqual == 0 {
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
	case useInt32 && !useInt64:
		return SERIES_TYPE_INT32, nil
	case hasInts:
		return SERIES_TYPE_INT64, nil
	default:
		return SERIES_TYPE_STRING, ErrCouldNotDetectType
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
	case "uint", "uint64", "uint32", "uint16", "uint8", "byte":
		return SERIES_TYPE_INT64, nil
	case "string":
		return SERIES_TYPE_STRING, nil
	case "bool":
		return SERIES_TYPE_BOOL, nil
	}
	return SERIES_TYPE_INVAILD, fmt.Errorf("type (%s) is not supported", s)
}

func detectTypes[T num.GenericType](v T) (Type, any) {
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
