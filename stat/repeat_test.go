package stat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	f32 := float32(1)
	f64 := float64(1)

	n := 10
	fs32 := Repeat(f32, n)
	fmt.Println(fs32)
	fs64 := Repeat(f64, n)
	fmt.Println(fs64)
}
