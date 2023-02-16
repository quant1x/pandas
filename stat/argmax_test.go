package stat

import (
	"fmt"
	"testing"
)

func TestArgMax(t *testing.T) {
	n1 := []float32{1.1, 2.2, 1.3, 1.4}
	n2 := []float64{1.2, 1.2, 3.3}
	n3 := []int64{11, 12, 33}
	fmt.Println(ArgMax2(n1))
	fmt.Println(ArgMax2(n2))
	fmt.Println(ArgMax2(n3))
}
