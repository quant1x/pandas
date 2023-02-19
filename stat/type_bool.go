package stat

import (
	"fmt"
	"github.com/mymmsc/gox/logger"
	"strconv"
)

const (
	Nil2Bool              = false      // 空指针转bool
	BoolNaN               = false      // bool 无效值
	True2Bool             = true       // true转bool
	False2Bool            = false      // false 转bool
	True2Float32  float32 = float32(1) // true转float32
	False2Float32 float32 = float32(0) // false转float32

	StringBad2Bool   = false // 字符串解析bool异常
	StringTrue2Bool  = true  // 字符串true转bool
	StringFalse2Bool = false // 字符串false转bool
)

func StringIsTrue(s string) bool {
	if s == "true" || s == "TRUE" || s == "True" || s == "1" || s == "真" || s == "对" || s == "好" {
		return true
	} else {
		return false
	}
}

func StringIsFalse(s string) bool {
	if s == "false" || s == "FALSE" || s == "False" || s == "0" || s == "假" || s == "错" || s == "坏" {
		return true
	} else {
		return false
	}
}

func BoolToInt(b bool) int8 {
	if b {
		return int8(1)
	}
	return int8(0)
}

func BoolToInt32(b bool) int32 {
	if b {
		return True2Int32
	}
	return False2Int32
}

func BoolToInt64(b bool) int64 {
	if b {
		return True2Int64
	}
	return False2Int64
}

// BoolToFloat32 bool转float32
func BoolToFloat32(b bool) float32 {
	if b {
		return True2Float32
	}
	return False2Float32
}

// BoolToFloat64 bool转float64
func BoolToFloat64(b bool) float64 {
	if b {
		return True2Float64
	}
	return False2Float64
}

// BoolToString bool 转 string
func BoolToString(b bool) string {
	if b {
		return True2String
	}
	return False2String
}

// ParseBool 字符串转bool
//
//	任意组合的nan字符串都会被解析成NaN
func ParseBool(s string, v any) bool {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("ParseBool %+v, error=%+v\n", v, err)
		}
	}()
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
		return integer2Bool[int](*val)
	case int:
		return integer2Bool[int](val)
	case *int64:
		if val == nil {
			return Nil2Bool
		}
		return integer2Bool[int64](*val)
	case int64:
		return integer2Bool[int64](val)
	case *float64:
		if val == nil {
			return Nil2Bool
		}
		return integer2Bool[float64](*val)
	case float64:
		return integer2Bool[float64](val)
	case *string:
		if val == nil {
			return Nil2Bool
		}
		if IsEmpty(*val) {
			return Nil2Bool
		}
		if StringIsTrue(*val) {
			return StringTrue2Bool
		} else if StringIsFalse(*val) {
			return StringFalse2Bool
		}
		f := ParseBool(*val, v)
		return f
	case string:
		if IsEmpty(val) {
			return Nil2Bool
		}
		if StringIsTrue(val) {
			return StringTrue2Bool
		} else if StringIsFalse(val) {
			return StringFalse2Bool
		}
		f := ParseBool(val, v)
		return f
	default:
		f := ParseBool(fmt.Sprintf("%v", v), v)
		return f
	}
}
