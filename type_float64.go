package pandas

import (
	"fmt"
	"github.com/mymmsc/gox/logger"
	"math"
	"strconv"
)

const (
	MaxFloat64          float64 = float64(math.MaxFloat64)             // float64最大值
	MinFloat64          float64 = float64(math.SmallestNonzeroFloat64) // float64最小值
	True2Float64        float64 = float64(1)                           // true转float64
	False2Float64       float64 = float64(0)                           // false转float64
	StringNil2Float     float64 = float64(0)                           // deprecated: 字符串空指针转float64
	StringBad2Float     float64 = float64(0)                           // deprecated: 字符串解析float64异常
	StringTrue2Float64  float64 = float64(1)                           // 字符串true转float64
	StringFalse2Float64 float64 = float64(0)                           // 字符串false转float64
)

// Float64IsNaN 判断float64是否NaN
func Float64IsNaN(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}

// ParseFloat64 字符串转float64
// 任意组合的nan字符串都会被解析成NaN
func ParseFloat64(s string, v any) float64 {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("ParseFloat64 %+v, error=%+v\n", v, err)
		}
	}()
	if IsEmpty(s) {
		return Nil2Float64
	}
	if isTrue(s) {
		return StringTrue2Float64
	} else if isFalse(s) {
		return StringFalse2Float64
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f
	}
	if IgnoreParseExceptions {
		return Nil2Float64
	}
	_ = v.(float64) // Intentionally panic
	return Nil2Float64
}

// float64转string
func float64ToString(v float64) string {
	if isNaN(v) {
		return StringNaN
	}
	return fmt.Sprintf("%f", v)
}

func AnyToFloat64(v any) float64 {
	if isPoint(v) {
		return pointToNumber[float64](v, Nil2Float64, boolToFloat64, ParseFloat64)
	}
	f := valueToNumber[float64](v, Nil2Float64, boolToFloat64, ParseFloat64)
	return f
}
