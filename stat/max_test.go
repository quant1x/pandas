package stat

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {

	f1 := []float32{1.1, 1.2, 1.3}
	f2 := []float32{1.1, 1.2, 1.3}
	fmt.Println(Max(f1))
	fmt.Println(Max(f2))

}
