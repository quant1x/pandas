package stat

import (
	"fmt"
	"testing"
)

func Test_PolyVal(t *testing.T) {
	x := []float64{0.0, 0.1, 0.2, 0.3, 0.5, 0.8, 1.0}
	y := []float64{1.0, 0.41, 0.50, 0.61, 0.91, 2.02, 2.46}
	A := PolyFit(x, y, 2)
	fmt.Println(A)

	//A2 := []float64{3.131561350718812, -1.2400367769976413, 0.7355767301905694}
	z1 := PolyVal(A, x)
	fmt.Println(z1)

	W := 5
	A2 := PolyFit(y, Sequence[float64](W), 1)
	x2 := Repeat[float64](float64(W), W)
	z2 := PolyVal(A2, x2)
	fmt.Println(z2)
}
