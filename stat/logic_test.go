package stat

import (
	"fmt"
	"testing"
)

func Test___compare(t *testing.T) {
	d1 := []float32{11, 12, 13, 14, 15}
	d2 := []float64{1, 2, 3, 34, 5}

	fmt.Println("----------<Gt>----------")
	fmt.Println(Gt(d1, d2))
	fmt.Println(Gt(d1, 13))
	fmt.Println("----------<Gte>----------")
	fmt.Println(Gte(d1, d2))
	fmt.Println(Gte(d1, 13))
	fmt.Println("----------<Lt>----------")
	fmt.Println(Lt(d1, d2))
	fmt.Println(Lt(d1, 13))
	fmt.Println("----------<Lte>----------")
	fmt.Println(Lte(d1, d2))
	fmt.Println(Lte(d1, 13))
}
