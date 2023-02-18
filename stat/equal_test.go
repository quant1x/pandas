package stat

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Equal[float64](d1, d2))
}
