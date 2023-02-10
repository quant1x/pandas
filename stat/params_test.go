package stat

import (
	"fmt"
	"testing"
)

func TestAnyToSlice(t *testing.T) {
	d1 := []float64{1, 2, 3, 4}
	fmt.Println(AnyToSlice[float64](float64(1), 5))
	fmt.Println(AnyToSlice[float64](d1, 3))
	fmt.Println(AnyToSlice[int32](d1, 5))
}
