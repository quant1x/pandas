package stat

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T) {
	fmt.Println(Median([]int8{}))
	d1 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d3 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d4 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Median(d1))
	fmt.Println(Median(d2))
	fmt.Println(Median(d3))
	fmt.Println(Median(d4))
}
