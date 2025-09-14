package formula

import (
	"fmt"
	"testing"

	"gitee.com/quant1x/pandas"
)

func TestHHVBARS(t *testing.T) {
	n1 := []float32{1.1, 2.2, 1.3, 1.4}
	s1 := pandas.NewSeries[float32](n1...)
	fmt.Println(HHVBARS(s1, 2))
}
