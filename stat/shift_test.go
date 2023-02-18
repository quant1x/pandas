package stat

import (
	"fmt"
	"testing"
)

func TestShift(t *testing.T) {
	d := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(ShiftN(d, 1))
	fmt.Println(ShiftN(d, -1))
	n1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(Shift(d, n1))
	n2 := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	fmt.Println(Shift(d, n2))

	n3 := []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	fmt.Println(Shift(d, n3))
	n4 := []int{1, -1, 1, -1, 1, -1, -1, -1, -1, 3}
	fmt.Println(Shift(d, n4))
}
