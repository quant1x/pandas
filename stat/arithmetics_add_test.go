package stat

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	f1 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f2 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d3 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d4 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Add(f1, f2))
	fmt.Println(Add(d2, float64(1)))
	fmt.Println(Add(d3, int32(2)))
	fmt.Println(Add(d4, int64(3)))
}
