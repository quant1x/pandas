package stat

import (
	"github.com/viterin/vek"
	"math"
)

var (
	// Nil2Float64 nil指针转换float64
	Nil2Float64 = float64(0)
	// Nil2Float32 nil指针转换float32
	Nil2Float32 = float32(0)
)

// 初始化 avx2
// 可以参考另一个实现库 gonum.org/v1/gonum/stat
func init() {
	// 开启加速选项
	vek.SetAcceleration(true)
	Nil2Float64 = math.NaN()
	// 这个转换是对的, NaN对float32也有效
	Nil2Float32 = float32(Nil2Float64)
}

// Float32IsNaN 判断float32是否NaN
func Float32IsNaN(f float32) bool {
	return Float64IsNaN(float64(f))
}

// Float64IsNaN 判断float64是否NaN
func Float64IsNaN(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}
