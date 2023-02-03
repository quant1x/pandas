package stat

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {

	v1 := []int32{1, -1, 2, -2}
	fmt.Println(Abs(v1))
	v2 := []int64{1, -1, 2, -2}
	fmt.Println(Abs(v2))
	v3 := []float32{1.1, -1.1, 2.2, -2.2}
	fmt.Println(Abs(v3))
	v4 := []float64{1.1, -1.1, 2.2, -2.2}
	fmt.Println(Abs(v4))

}
