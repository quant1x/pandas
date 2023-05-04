package stat

import (
	"fmt"
	"github.com/mymmsc/gox/logger"
	"math"
	"strconv"
)

const (
	MaxInt32          = int32(math.MaxInt32)
	MinInt32          = int32(math.MinInt32)
	Nil2Int32         = int32(0) // 空指针转int32
	Int32NaN          = int32(0) // int32 无效值
	True2Int32        = int32(1) // true转int32
	False2Int32       = int32(0) // false 转int32
	StringBad2Int32   = int32(0) // 字符串解析int32异常
	StringTrue2Int32  = int32(1) // 字符串true转int32
	StringFalse2Int32 = int32(0) // 字符串false转int32
)

// ParseInt32 解析int字符串, 尝试解析10进制和16进制
func ParseInt32(s string, v any) int32 {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("ParseInt32 %+v, error=%+v\n", v, err)
		}
	}()
	if IsEmpty(s) {
		return Nil2Int32
	}
	if StringIsTrue(s) {
		return StringTrue2Int32
	} else if StringIsFalse(s) {
		return StringFalse2Int32
	}
	if StringIsNaN(s) {
		return StringFalse2Int32
	}
	i, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return int32(i)
	}
	// 解析失败继续解析16进制
	i, err = strconv.ParseInt(s, 16, 32)
	if err == nil {
		return int32(i)
	}
	logger.Errorf("%s, error=%+v\n", s, err)
	if IgnoreParseExceptions {
		i = int64(StringBad2Int32)
	} else {
		_ = v.(int32) // Intentionally panic
	}
	return int32(i)
}

func int32ToString(v int32) string {
	if Float64IsNaN(float64(v)) {
		return StringNaN
	}
	return fmt.Sprint(v)
}

// AnyToInt32 any转换int32
func AnyToInt32(v any) int32 {
	if vv, ok := extraceValueFromPointer(v); ok {
		v = vv
	}

	f := valueToNumber[int32](v, Nil2Int32, BoolToInt32, ParseInt32)
	return f
}
