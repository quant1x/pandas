package stat

import (
	"fmt"
	"testing"
)

func TestNewFrameT1(t *testing.T) {
	f1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, NaN(), 12}
	f2 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s1 := []string{"a", "b", "c"}
	i1 := []int64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	gf1 := NewFrame[float64]("x", f1)
	fmt.Printf("%+v\n", gf1)

	t0 := []any{nil, 1, true, "abc", 3.45, NaN()}
	gf2 := NewFrame[float64]("x", t0...)
	fmt.Printf("%+v\n", gf2)
	_ = f1
	_ = f2
	_ = s1
	_ = i1
	_ = gf1
}
