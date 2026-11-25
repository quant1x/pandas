package formula

import (
	"fmt"
	"math"
	"testing"

	"github.com/quant1x/pandas"
)

func TestLAST(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12}
	i0 := CompareGt(f0, 10)
	s0 := pandas.NewSeriesWithoutType("f0", i0)
	fmt.Println(LAST(s0, 10, 5))
}

// TestLASTWindowShortage 锁定窗口不足的返回口径：NaN，而不是「视为成立」。
func TestLASTWindowShortage(t *testing.T) {
	condition := pandas.Vector([]float64{1, 1, 1, 0, 1, 1})
	want := []float64{math.NaN(), math.NaN(), math.NaN(), 1, 0, 0}
	got := LAST(condition, 3, 1).DTypes()
	if len(got) != len(want) {
		t.Fatalf("长度 %d != %d", len(got), len(want))
	}
	for i := range want {
		if math.IsNaN(want[i]) && math.IsNaN(got[i]) {
			continue
		}
		if got[i] != want[i] {
			t.Fatalf("第 %d 位 %v != %v", i, got[i], want[i])
		}
	}
}
