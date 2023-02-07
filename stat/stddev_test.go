package stat

import (
	"fmt"
	"testing"
)

func TestStd(t *testing.T) {
	f0 := []float64{1.1, 2.2, 1.3, 1.4}
	f1 := []float64{70, 80, 75, 83, 86}
	f2 := []float64{90, 69, 60, 88, 87}
	fmt.Println(Std(f0))
	fmt.Println(Std(f1))
	fmt.Println(Std(f2))
}
