package stat

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	d1 := []bool{true, true}
	d2 := []bool{true, false}
	d3 := []uintptr{0, 1}
	d4 := []int{1, 1}
	d5 := []float64{1.0, 0}
	fmt.Println(Count(d1))
	fmt.Println(Count(d2))
	fmt.Println(Count(d3))
	fmt.Println(Count(d4))
	fmt.Println(Count(d5))
}
