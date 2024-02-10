package stat

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	v1 := []int32{1, -1, 2, -2}
	fmt.Println(Abs(v1))
	v2 := []int64{1, -1, 2, -2}
	fmt.Println(Abs(v2))
	v3 := []float32{1.1, -1.1, 2.2, -2.2}
	fmt.Println(Abs(v3))
	v4 := []float64{1.1, -1.1, 2.2, -2.2}
	fmt.Println(Abs(v4))

	v5 := []uint{1, 2, 3, 4}
	fmt.Println(Abs(v5))
	v6 := []int8{1, -1, 2, -2}
	fmt.Println(Abs(v6))
	v7 := []int16{1, -1, 2, -2}
	fmt.Println(Abs(v7))
	v8 := []int64{1, -1, 2, -2}
	fmt.Println(Abs(v8))
	v9 := []int{1, -1, 2, -2}
	fmt.Println(Abs(v9))
}
