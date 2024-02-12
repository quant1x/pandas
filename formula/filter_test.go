package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestFILTER(t *testing.T) {
	n1 := []float32{1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(n1)
	s1 := pandas.NewNDArray[float32](n1...)
	fmt.Println(FILTER(s1, 5))

	//w := 2
	//for i := 0; i < len(n1); i++ {
	//	if n1[i] != 0 {
	//		for j := i + 1; j < i+1+w; j++ {
	//			n1[j] = 0
	//		}
	//	}
	//}
	//fmt.Println(n1)
}
