package pandas

import (
	"fmt"
	"github.com/mymmsc/gox/logger"
	"strconv"
)

const (
	Nil2Int         = int64(0) // 空指针转int64
	IntNaN          = int64(0) // int64 无效值
	True2Int        = int64(1) // true转int64
	False2Int       = int64(0) // false 转int64
	StringBad2Int   = int64(0) // 字符串解析int64异常
	StringTrue2Int  = int64(1) // 字符串true转int64
	StringFalse2Int = int64(0) // 字符串false转int64
)

// AnyToInt64 any转换int64
func AnyToInt64(v any) int64 {
	switch val := v.(type) {
	case nil:
		return IntNaN
	case *bool:
		if val == nil {
			return Nil2Int
		}
		if *val == true {
			return True2Int
		}
		return False2Int
	case bool:
		if val == true {
			return True2Int
		}
		return False2Int
	case *int:
		if val == nil {
			return Nil2Int
		}
		return int64(*val)
	case int:
		return int64(val)
	case *int64:
		if val == nil {
			return Nil2Int
		}
		return *val
	case int64:
		return val
	case *float32:
		if val == nil {
			return Nil2Int
		}
		return int64(*val)
	case float32:
		return int64(val)
	case *float64:
		if val == nil {
			return Nil2Int
		}
		return int64(*val)
	case float64:
		return int64(val)
	case *string:
		if val == nil {
			return Nil2Int
		}
		if IsEmpty(*val) {
			return Nil2Int
		}
		if isTrue(*val) {
			return StringTrue2Int
		} else if isFalse(*val) {
			return StringFalse2Int
		}
		i := ParseInt(*val, v)
		return i
	case string:
		if IsEmpty(val) {
			return Nil2Int
		}
		if isTrue(val) {
			return StringTrue2Int
		} else if isFalse(val) {
			return StringFalse2Int
		}
		i := ParseInt(val, v)
		return i
	default:
		i := ParseInt(fmt.Sprintf("%v", v), v)
		return i
	}
}

// ParseInt 解析int字符串, 尝试解析10进制和16进制
func ParseInt(s string, v any) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i, err = strconv.ParseInt(s, 16, 64)
		if err != nil {
			logger.Errorf("%s, error=%+v\n", s, err)
			if IgnoreParseExceptions {
				i = StringBad2Int
			} else {
				_ = v.(int64) // Intentionally panic
			}
		}
	}
	return i
}

func int2String(v int64) string {
	if isNaN(float64(v)) {
		return StringNaN
	}
	return fmt.Sprint(v)
}
