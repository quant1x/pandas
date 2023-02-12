package stat

import (
	"fmt"
	"testing"
)

func TestCumSum(t *testing.T) {
	n1 := []float32{1.1, 2.2, 1.3, 1.4}
	n2 := []float64{1.2, 1.2, 3.3}
	n3 := []int8{1, 2, 3}
	fmt.Println(CumSum(n1))
	fmt.Println(CumSum(n2))
	fmt.Println(CumSum(n3))
}
