package stat

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	f0 := []float64{1.1, 2.2, 1.3, 1.4}
	f1 := []float64{70, 80, 75, 83, 86}
	f2 := []float64{90, 69, 60, 88, 87}
	r0 := f0[0:1:1]
	fmt.Println(Reverse(f0))
	fmt.Println(Reverse(f1))
	fmt.Println(Reverse(f2))

	_ = r0
}
