package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestLLVBARS(t *testing.T) {
	n1 := []float32{1.1, 2.2, 1.3, 1.4}
	s1 := stat.NewNDArray[float32](n1...)
	fmt.Println(LLVBARS(s1, 2))
}
