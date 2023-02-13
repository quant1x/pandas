package stat

import (
	"fmt"
	"testing"
)

func TestMul(t *testing.T) {
	f1 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f2 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d3 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d4 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Mul(f1, f2))
	fmt.Println(Mul(d2, float64(1)))
	fmt.Println(Mul(d3, int32(2)))
	fmt.Println(Mul(d4, int64(3)))
	fmt.Println(Mul(d4, int64(3)))
}
