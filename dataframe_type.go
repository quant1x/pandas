package pandas

import (
	"fmt"
	"strconv"
)

func findTypeByString(arr []string) (Type, error) {
	var hasFloats, hasInts, hasBools, hasStrings bool
	for _, str := range arr {
		if str == "" || str == "NaN" {
			continue
		}
		if _, err := strconv.Atoi(str); err == nil {
			hasInts = true
			continue
		}
		if _, err := strconv.ParseFloat(str, 64); err == nil {
			hasFloats = true
			continue
		}
		if str == "true" || str == "false" {
			hasBools = true
			continue
		}
		hasStrings = true
	}
	// 类型优先级, string > bool > float > int, string 为默认类型
	switch {
	case hasStrings:
		return SERIES_TYPE_STRING, nil
	case hasBools:
		return SERIES_TYPE_BOOL, nil
	case hasFloats:
		return SERIES_TYPE_FLOAT, nil
	case hasInts:
		return SERIES_TYPE_INT, nil
	default:
		return SERIES_TYPE_STRING, fmt.Errorf("couldn't detect type")
	}
}

// 检测类型
func detectTypeBySlice(arr []any) (Type, error) {
	var hasFloats, hasInts, hasBools, hasStrings bool
	for _, v := range arr {
		switch value := v.(type) {
		case string:
			hasStrings = true
			continue
		case float32, float64:
			hasFloats = true
			continue
		case int, int32, int64:
			hasInts = true
			continue
		case bool:
			hasBools = true
			continue
		default:
			_ = value
		}
	}

	switch {
	case hasStrings:
		return SERIES_TYPE_STRING, nil
	case hasBools:
		return SERIES_TYPE_BOOL, nil
	case hasFloats:
		return SERIES_TYPE_FLOAT, nil
	case hasInts:
		return SERIES_TYPE_INT, nil
	default:
		return SERIES_TYPE_STRING, fmt.Errorf("couldn't detect type")
	}
}
