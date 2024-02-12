package stat

import (
	"fmt"
	"gitee.com/quant1x/num"
	"testing"
)

func TestNewSeries(t *testing.T) {
	d1 := []num.DType{}
	fmt.Println(d1)
	s1 := NewSeries[num.DType]()
	fmt.Println(s1)
}
