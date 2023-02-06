package stat

import (
	"fmt"
	"testing"
)

func TestMaxinum(t *testing.T) {
	f1 := []float32{1.1, 2.2, 1.3, 1.4}
	f2 := []float32{1.2, 1.2, 3.3}
	fmt.Println(Maxinum(f1, f2))
}
