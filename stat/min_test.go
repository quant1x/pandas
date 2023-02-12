package stat

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	f1 := []float32{1.1, 1.2, 1.3}
	f2 := []float64{1.1, 1.2, 1.3}
	i1 := []int8{1, 2, 3, -1}
	fmt.Println(Min(f1))
	fmt.Println(Min(f2))
	fmt.Println(Min(i1))
}
