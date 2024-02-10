package stat

import (
	"fmt"
	"gitee.com/quant1x/gox/logger"
	"math"
	"strconv"
)

const (
	MaxInt64          = int64(math.MaxInt64)
	MinInt64          = int64(math.MinInt64)
	Nil2Int64         = int64(0) // 空指针转int64
	Int64NaN          = int64(0) // int64 无效值
	True2Int64        = int64(1) // true转int64
	False2Int64       = int64(0) // false 转int64
	StringBad2Int64   = int64(0) // 字符串解析int64异常
	StringTrue2Int64  = int64(1) // 字符串true转int64
	StringFalse2Int64 = int64(0) // 字符串false转int64
)

// ParseInt64 解析int字符串, 尝试解析10进制和16进制
func ParseInt64(s string, v any) int64 {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("ParseInt64 %+v, error=%+v\n", v, err)
		}
	}()
	if IsEmpty(s) {
		return Nil2Int64
	}
	if StringIsTrue(s) {
		return StringTrue2Int64
	} else if StringIsFalse(s) {
		return StringFalse2Int64
	}
	if StringIsNaN(s) {
		return StringFalse2Int64
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}
	// 解析失败继续解析16进制
	i, err = strconv.ParseInt(s, 16, 64)
	if err == nil {
		return i
	}
	logger.Errorf("%s, error=%+v\n", s, err)
	if IgnoreParseExceptions {
		i = StringBad2Int64
	} else {
		_ = v.(int64) // Intentionally panic
	}
	return i
}

func int64ToString(v int64) string {
	if Float64IsNaN(float64(v)) {
		return StringNaN
	}
	return fmt.Sprint(v)
}

// AnyToInt64 any转换int64
func AnyToInt64(v any) int64 {
	if vv, ok := ExtractValueFromPointer(v); ok {
		v = vv
	}

	f := valueToNumber[int64](v, Nil2Int64, BoolToInt64, ParseInt64)
	return f
}
