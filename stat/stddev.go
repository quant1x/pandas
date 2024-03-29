package stat

import (
	"gitee.com/quant1x/vek"
	"gitee.com/quant1x/vek/vek32"
	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/stat"
	"math"
)

// Std_TODO StdDev 这个版本有bug, gonum计算的std不对
// TODO: 于总来解决
func Std_TODO[T Float](f []T) T {
	if len(f) == 0 {
		return T(0)
	}
	var d any
	var s any
	s = f
	switch fs := s.(type) {
	case []float32:
		d = f32_std(fs)
	case []float64:
		// 这里计算不对
		d = stat.StdDev(fs, nil)
	default:
		// 应该不会走到这里
		panic(ErrUnsupportedType)
	}

	return d.(T)
}

// Std 计算标准差
func Std[T BaseType](f []T) T {
	if len(f) == 0 {
		return typeDefault[T]()
	}
	var d any
	var s any
	s = f
	switch fs := s.(type) {
	case []float32:
		d = f32_std(fs)
	case []float64:
		d = f64_std(fs)
	default:
		// 应该不会走到这里
		panic(ErrUnsupportedType)
	}

	return d.(T)
}

func f64_std(f []float64) float64 {
	values := slices.Clone(f)
	// 求平均数
	meam := vek.Mean(values)
	// 减去 平均数
	vek.SubNumber_Inplace(values, meam)
	// 计算方差
	y := vek.Repeat(2.00, len(f))
	vek.Pow_Inplace(values, y)
	// 再求方差平均数
	meam = vek.Mean(values)
	meam = math.Sqrt(meam)
	return meam
}

func f32_std(f []float32) float32 {
	values := slices.Clone(f)
	// 求平均数
	meam := vek32.Mean(values)
	// 减去 平均数
	vek32.SubNumber_Inplace(values, meam)
	// 计算方差
	y := vek32.Repeat(2.00, len(f))
	vek32.Pow_Inplace(values, y)
	// 再求方差平均数
	meam = vek32.Mean(values)
	meam = float32(math.Sqrt(float64(meam)))
	return meam
}
