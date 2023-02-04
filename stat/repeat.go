package stat

import (
	"github.com/viterin/vek"
	"github.com/viterin/vek/vek32"
	"unsafe"
)

// Repeat repeat
func Repeat[T Float](f T, n int) []T {
	var d any
	bitsize := unsafe.Sizeof(f)
	if bitsize == 4 {
		d = vek32.Repeat(float32(f), n)
	} else if bitsize == 8 {
		d = vek.Repeat(float64(f), n)
	} else {
		// 应该不会走到这里
		d = []T{}
	}
	return d.([]T)
}
