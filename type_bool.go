package pandas

import (
	"fmt"
	"strconv"
)

const (
	Nil2Bool         = false // 空指针转int64
	BoolNaN          = false // int64 无效值
	True2Bool        = true  // true转int64
	False2Bool       = false // false 转int64
	StringBad2Bool   = false // 字符串解析int64异常
	StringTrue2Bool  = true  // 字符串true转int64
	StringFalse2Bool = false // 字符串false转int64
)

// AnyToBool any转换bool
func AnyToBool(v any) bool {
	switch val := v.(type) {
	case nil:
		return Nil2Bool
	case *bool:
		if val == nil {
			return Nil2Bool
		}
		if *val == true {
			return True2Bool
		}
		return False2Bool
	case bool:
		if val == true {
			return True2Bool
		}
		return False2Bool
	case *int:
		if val == nil {
			return Nil2Bool
		}
		return int2Bool(*val)
	case int:
		return int2Bool(val)
	case *int64:
		if val == nil {
			return Nil2Bool
		}
		return int2Bool(*val)
	case int64:
		return int2Bool(val)
	case *float64:
		if val == nil {
			return Nil2Bool
		}
		return float2Bool(*val)
	case float64:
		return float2Bool(val)
	case *string:
		if val == nil {
			return Nil2Bool
		}
		if IsEmpty(*val) {
			return Nil2Bool
		}
		if isTrue(*val) {
			return StringTrue2Bool
		} else if isFalse(*val) {
			return StringFalse2Bool
		}
		f := ParseBool(*val, v)
		return f
	case string:
		if IsEmpty(val) {
			return Nil2Bool
		}
		if isTrue(val) {
			return StringTrue2Bool
		} else if isFalse(val) {
			return StringFalse2Bool
		}
		f := ParseBool(val, v)
		return f
	default:
		f := ParseBool(fmt.Sprintf("%v", v), v)
		return f
	}
}

// ParseBool 字符串转bool
// 任意组合的nan字符串都会被解析成NaN
func ParseBool(s string, v any) bool {
	f, err := strconv.ParseBool(s)
	if err != nil {
		if IgnoreParseExceptions {
			f = Nil2Bool
		} else {
			_ = v.(float64) // Intentionally panic
		}
	}
	return f
}

func isTrue(s string) bool {
	if s == "true" || s == "TRUE" || s == "True" || s == "1" || s == "真" || s == "对" || s == "好" {
		return true
	} else {
		return false
	}
}

func isFalse(s string) bool {
	if s == "false" || s == "FALSE" || s == "False" || s == "0" || s == "假" || s == "错" || s == "坏" {
		return true
	} else {
		return false
	}
}

func bool2String(v bool) string {
	if v {
		return True2String
	}
	return False2String
}

func int2Bool[T ~int | ~int32 | int64](n T) bool {
	if n == 0 {
		return false
	}
	return true
}

func float2Bool[T ~float32 | float64](f T) bool {
	if IsNaN(float64(f)) || f == 0 {
		return false
	}
	return true
}
