package stat

import (
	"fmt"
	"testing"
)

func TestAnyToSlice(t *testing.T) {
	d1 := []float64{1, 2, 3, 4}
	d2 := []int32{1, 2, 3, 4}
	d3 := []bool{true, true}
	d4 := []string{"a", "b"}
	fmt.Println(AnyToSlice[float64](float64(1), 5))
	fmt.Println(AnyToSlice[float64](d1, 3))
	fmt.Println(AnyToSlice[int32](d2, 5))
	fmt.Println(AnyToSlice[bool](d3, 5))
	fmt.Println(AnyToSlice[string](d4, 5))
	fmt.Println(AnyToSlice[string](nil, 5))
	fmt.Println(AnyToSlice[string]([]string{"a"}, 5))
	fmt.Println(AnyToSlice[string]("a", 5))
	fmt.Println(AnyToSlice[bool](true, 5))
	fmt.Println(AnyToSlice[bool]([]bool{true}, 5))
}
