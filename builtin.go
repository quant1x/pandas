package pandas

import (
	"math"
	"strings"
)

// 收敛统一初始化

// 全局变量定义

var (
	// Nil2Float nil指针转换float64
	Nil2Float = float64(0)
)

func init() {
	Nil2Float = math.NaN()
}

// NaN returns an IEEE 754 “not-a-number” value.
func NaN() float64 {
	return math.NaN()
}

// IsNan float64是否NaN
func IsNaN(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}

// IsEmpty Code to test if string is empty
func IsEmpty(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	} else {
		return false
	}
}

// Repeat 重复生成a
func Repeat[T SeriesGenericType](a T, n int) []T {
	dst := make([]T, n)
	for i := 0; i < n; i++ {
		dst[i] = a
	}
	return dst
}
