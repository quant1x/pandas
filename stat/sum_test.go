package stat

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	d1 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d3 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d4 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Sum(d1))
	fmt.Println(Sum(d2))
	fmt.Println(Sum(d3))
	fmt.Println(Sum(d4))
}
