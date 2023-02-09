package stat

import (
	"fmt"
	"testing"
)

func TestFill(t *testing.T) {
	s2 := []float64{1, 2, 3, 4, 3, 3, 2, 1, DTypeNaN, DTypeNaN, DTypeNaN, DTypeNaN}
	fmt.Println(s2)
	Fill(s2, 1.0)
	fmt.Println(s2)
	Fill(s2, 1.0, true)
	fmt.Println(s2)
}
