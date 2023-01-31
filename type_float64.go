package pandas

import (
	"fmt"
	"strconv"
)

const (
	True2Float        float64 = float64(1) // true转float64
	False2Float       float64 = float64(0)
	StringNil2Float   float64 = float64(0) // deprecated: 字符串空指针转float64
	StringBad2Float   float64 = float64(0) // deprecated: 字符串解析float64异常
	StringTrue2Float  float64 = float64(1) // 字符串true转float64
	StringFalse2Float float64 = float64(0) // 字符串false转float64
)

// AnyToFloat64 any转换float64
func AnyToFloat64(v any) float64 {
	switch val := v.(type) {
	case nil:
		return nan()
	case *bool:
		if val == nil {
			return Nil2Float
		}
		if *val == true {
			return True2Float
		}
		return False2Float
	case bool:
		if val == true {
			return True2Float
		}
		return False2Float
	case *int:
		if val == nil {
			return Nil2Float
		}
		return float64(*val)
	case int:
		return float64(val)
	case *int64:
		if val == nil {
			return Nil2Float
		}
		return float64(*val)
	case int64:
		return float64(val)
	case *float64:
		if val == nil {
			return Nil2Float
		}
		return *val
	case float64:
		return val
	case *string:
		if val == nil {
			return Nil2Float
		}
		if IsEmpty(*val) {
			return Nil2Float
		}
		if isTrue(*val) {
			return StringTrue2Float
		} else if isFalse(*val) {
			return StringFalse2Float
		}
		f := ParseFloat(*val, v)
		return f
	case string:
		if IsEmpty(val) {
			return Nil2Float
		}
		if isTrue(val) {
			return StringTrue2Float
		} else if isFalse(val) {
			return StringFalse2Float
		}
		f := ParseFloat(val, v)
		return f
	default:
		f := ParseFloat(fmt.Sprintf("%v", v), v)
		return f
	}
}

// ParseFloat 字符串转float64
// 任意组合的nan字符串都会被解析成NaN
func ParseFloat(s string, v any) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		if IgnoreParseExceptions {
			f = Nil2Float
		} else {
			_ = v.(float64) // Intentionally panic
		}
	}
	return f
}

func float2String(v float64) string {
	if isNaN(v) {
		return StringNaN
	}
	return fmt.Sprintf("%f", v)
}
