package stat

import (
	"fmt"
	"testing"
)

func TestAny(t *testing.T) {
	d1 := []bool{true, true}
	d2 := []bool{true, false}
	d3 := []uintptr{0, 1}
	d4 := []int{1, 1}
	d5 := []float64{1.0, 0}
	fmt.Println(Any(d1))
	fmt.Println(Any(d2))
	fmt.Println(Any(d3))
	fmt.Println(Any(d4))
	fmt.Println(Any(d5))
}
