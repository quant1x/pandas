package stat

import (
	"fmt"
	"testing"
)

func TestMinimum(t *testing.T) {
	f1 := []float32{1.1, 2.2, 1.3, 1.4}
	f2 := []float32{1.2, 1.2, 3.3}
	fmt.Println(Minimum(f1, f2))
	fmt.Println(Minimum([]float64{1.0}, []float64{1.0}))
	fmt.Println(Minimum([]int{1}, []int{1}))
	fmt.Println(Minimum([]int{1}, []int{2}))
	fmt.Println(Minimum([]int{2}, []int{1, 2, 3}))
	fmt.Println(Minimum([]int{1, 2, 3, 4, 5}, []int{2}))
}
