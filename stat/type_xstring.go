package stat

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

// AnyToString any转string
func AnyToString(v any) string {
	if vv, ok := extraceValueFromPointer(v); ok {
		v = vv
	}

	s := __anyToString(v)
	return s
}
