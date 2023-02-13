package stat

import (
	"fmt"
	"testing"
)

func TestDiv(t *testing.T) {
	f1 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f2 := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d3 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d4 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Div(f1, f2))
	fmt.Println(Div(d2, float64(1)))
	fmt.Println(Div(d3, int32(2)))
	fmt.Println(Div(d4, int64(3)))
	fmt.Println(Div(d4, int64(3)))
}
