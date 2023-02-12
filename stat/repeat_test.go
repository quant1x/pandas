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

func TestSequence(t *testing.T) {
	fmt.Println(Range[float32](5))
	fmt.Println(Range[float64](5))
	fmt.Println(Range[int8](5))
	fmt.Println(Range[uint8](5))
	fmt.Println(Range[int16](5))
	fmt.Println(Range[uint16](5))
	fmt.Println(Range[int32](5))
	fmt.Println(Range[uint32](5))
	fmt.Println(Range[int64](5))
	fmt.Println(Range[uint64](5))
	fmt.Println(Range[int](5))
	fmt.Println(Range[uint](5))
	fmt.Println(Range[uintptr](5))
}
