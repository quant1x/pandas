package pandas

import (
	gc "github.com/huandu/go-clone"
	"math"
	"reflect"
	"strings"
)

// 收敛统一初始化
const (
	quant1xPath = "~/.quant1x"         // quant1x默认
	tmpDir      = quant1xPath + "/tmp" // 临时路径
)

// 全局变量定义

var (
	// Nil2Float64 nil指针转换float64
	Nil2Float64 = float64(0)
	// Nil2Float32 nil指针转换float32
	Nil2Float32 = float32(0)
)

func init() {
	Nil2Float64 = math.NaN()
	// 这个转换是对的, NaN对float32也有效
	Nil2Float32 = float32(Nil2Float64)
}

// NaN returns an IEEE 754 “not-a-number” value.
func NaN() float64 {
	return math.NaN()
}

// IsNaN float64是否NaN
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

// Clone 克隆一个any
func clone(v any) any {
	return gc.Clone(v)
}

func isPoint(v any) bool {
	kind := reflect.ValueOf(v).Kind()
	return reflect.Pointer == kind
}
