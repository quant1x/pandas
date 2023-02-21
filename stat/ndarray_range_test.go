package stat

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNDArray_IndexOf(t *testing.T) {
	d1 := []string{"a0", "a1", "a2", "a3", "a4"}
	s1 := NDArray[string](d1)
	fmt.Println(s1)
	v1 := s1.IndexOf(1, true)
	if mv, ok := v1.(reflect.Value); ok {
		mv.SetString("1")
		fmt.Println(s1)
	}
}
