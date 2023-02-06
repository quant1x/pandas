package stat

import (
	"fmt"
	"math"
	"testing"
)

func TestMaxinum(t *testing.T) {
	fmt.Println(1.4 > math.NaN())
	fmt.Println(1.4 < math.NaN())
	f1 := []float32{1.1, 2.2, 1.3, 1.4}
	f2 := []float32{1.2, 1.2, 3.3}
	fmt.Println(MaximumAvx2(f1, f2))
	fmt.Println(Maximum(f1, f2))
}
