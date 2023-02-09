package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
)

// Sqrt 求平方根
func Sqrt[T StatType](v []T) []T {
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
