package stat

import (
	"fmt"
	"testing"
)

func TestNewSeries(t *testing.T) {
	d1 := []DType{}
	fmt.Println(d1)
	s1 := NewSeries[DType]()
	fmt.Println(s1)
}
