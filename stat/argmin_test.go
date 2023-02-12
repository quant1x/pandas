package stat

import (
	"fmt"
	"testing"
)

func TestArgMin(t *testing.T) {
	n1 := []float32{1.1, 2.2, 1.3, 1.4}
	n2 := []float64{1.2, 1.2, 0.3}
	n3 := []int64{11, 12, 33}
	n4 := []int32{55, 11, 12, 33}
	fmt.Println(ArgMin(n1))
	fmt.Println(ArgMin(n2))
	fmt.Println(ArgMin(n3))
	fmt.Println(ArgMin(n4))
}
