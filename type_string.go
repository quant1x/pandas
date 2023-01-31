package pandas

import (
	"github.com/mymmsc/gox/logger"
	"strconv"
	"strings"
)

const (
	StringNaN    = "NaN"   // 字符串NaN
	Nil2String   = "NaN"   // nil指针转string
	True2String  = "true"  // true转string
	False2String = "false" // false转string
)

// AnyToString any转string
func AnyToString(v any) string {
	switch val := v.(type) {
	case nil:
		return Nil2String
	case *bool:
		if val == nil {
			return Nil2String
		}
		if *val == true {
			return True2String
		} else {
			return False2String
		}
	case bool:
		if val == true {
			return True2String
		} else {
			return False2String
		}
	case *string:
		if val == nil {
			return Nil2String
		}
		return []string{*val}[0]
	case string:
		return val
	case *float64:
		if val == nil {
			return Nil2String
		}
		return []string{strconv.FormatFloat(*val, 'G', -1, 64)}[0]
	case float64:
		return []string{strconv.FormatFloat(val, 'G', -1, 64)}[0]
	case *float32:
		if val == nil {
			return Nil2String
		}
		return []string{strconv.FormatFloat(float64(*val), 'G', -1, 64)}[0]
	case float32:
		return []string{strconv.FormatFloat(float64(val), 'G', -1, 64)}[0]
	case *int64:
		if val == nil {
			return Nil2String
		}
		return []string{strconv.FormatInt(*val, 10)}[0]
	case int64:
		return []string{strconv.FormatInt(val, 10)}[0]
	case *int:
		if val == nil {
			return Nil2String
		}
		return []string{strconv.Itoa(*val)}[0]
	case int:
		return []string{strconv.Itoa(val)}[0]
	case *int32:
		if val == nil {
			return Nil2String
		}
		return []string{strconv.FormatInt(int64(*val), 10)}[0]
	case int32:
		return []string{strconv.FormatInt(int64(val), 10)}[0]
	default:
		logger.Errorf("%s, error=The type is not recognized\n", v)
		_ = v.(string) // Intentionally panic
		return Nil2String
	}
}

// StringIsNaN 判断字符串是否NaN
func StringIsNaN(s string) bool {
	s = strings.TrimSpace(s)
	if strings.ToLower(s) == "nan" {
		return true
	}
	return false
}
