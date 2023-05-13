package stat

import (
	"gitee.com/quant1x/gox/vek"
	"gitee.com/quant1x/gox/vek/vek32"
	"github.com/chewxy/math32"
	"math"
)

// Sqrt 求平方根
func Sqrt[T Number](v []T) []T {
	var d any
	var values any = v
	switch fs := values.(type) {
	case []float32:
		d = vek32.Sqrt(fs)
	case []float64:
		d = vek.Sqrt(fs)
	default:
		panic(ErrUnsupportedType)
	}

	return d.([]T)
}

func __sqrt_go_f64(x []float64) {
	for i := 0; i < len(x); i++ {
		x[i] = math.Sqrt(x[i])
	}
}

func __sqrt_go_f32(x []float32) {
	for i := 0; i < len(x); i++ {
		x[i] = math32.Sqrt(x[i])
	}
}
