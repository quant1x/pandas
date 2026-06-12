package formula

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/quant1x/pandas"
)

func TestCOUNT(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 11, 12}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&f0))
	fmt.Println(sh)
	i0 := CompareGte(f0, 1)
	fmt.Println(i0)
	s0 := pandas.NewSeriesWithoutType("f0", i0).Values().([]bool)
	fmt.Println(s0)
	fmt.Println(COUNT(s0, 5))
}
