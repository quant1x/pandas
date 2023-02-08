package stat

import (
	"fmt"
	"testing"
)

func TestArray_Len(t *testing.T) {
	f1 := []float64{1, 2, 3, 4, 5}
	a1 := Array[float64](f1)
	fmt.Println(a1)
	fmt.Println(a1.Len())
}
