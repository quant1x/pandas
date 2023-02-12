package stat

import (
	"fmt"
	"testing"
)

func TestZeros(t *testing.T) {
	r1 := Zeros[int](5)
	fmt.Println(r1)
}
