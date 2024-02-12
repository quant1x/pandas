package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestHHVBARS(t *testing.T) {
	n1 := []float32{1.1, 2.2, 1.3, 1.4}
	s1 := pandas.NewNDArray[float32](n1...)
	fmt.Println(HHVBARS(s1, 2))
}
