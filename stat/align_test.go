package stat

import (
	"fmt"
	"testing"
)

func TestAlign(t *testing.T) {

}

func TestAlign2Series(t *testing.T) {
	s := Align2Series(1, 10)
	fmt.Println(s)

	s1 := Align2Series([]int64{1, 2, 1}, 10)
	fmt.Println(s1)

	s2 := Align2Series(NDArray[int64]([]int64{1, 2, 1}), 10)
	fmt.Println(s2)
}
