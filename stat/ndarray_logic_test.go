package stat

import (
	"fmt"
	"testing"
)

func TestNDArray_Gt(t *testing.T) {
	d1 := []float32{11, 12, 13, 14, 15}
	d2 := []float64{1, 2, 3, 34, 5}

	s1 := NDArray[float32](d1)
	s2 := NDArray[float64](d2)

	fmt.Println("----------<Gt>----------")
	fmt.Println(s1.Gt(s2))
	fmt.Println(s1.Gt(13))
	fmt.Println("----------<Gte>----------")
	fmt.Println(Gte(d1, d2))
	fmt.Println(Gte(d1, 13))
	fmt.Println("----------<Lt>----------")
	fmt.Println(Lt(d1, d2))
	fmt.Println(Lt(d1, 13))
	fmt.Println("----------<Lte>----------")
	fmt.Println(Lte(d1, d2))
	fmt.Println(And(d1, 13))
	fmt.Println(Or(d1, 13))
}
