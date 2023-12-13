package stat

import (
	"fmt"
	"testing"
)

func Test___compare(t *testing.T) {
	d1 := []float32{11, 12, 13, 14, 15}
	d2 := []float64{1, 2, 3, 34, 5}
	d3 := []float64{1, 2, 3, 34}

	fmt.Println("----------<Gt>----------")
	fmt.Println(Gt(d1, d2))
	fmt.Println(Gt(d1, d3))
	fmt.Println(Gt(d2, d3))
	fmt.Println(Gt(d1, 13))
	fmt.Println("----------<Gte>----------")
	fmt.Println(Gte(d1, d2))
	fmt.Println(Gte(d1, d3))
	fmt.Println(Gte(d2, d3))
	fmt.Println(Gte(d1, 13))
	fmt.Println("----------<Lt>----------")
	fmt.Println(Lt(d1, d2))
	fmt.Println(Lt(d1, 13))
	fmt.Println("----------<Lte>----------")
	l1 := Lte(d1, d2)
	l2 := Lte(d1, 13)
	fmt.Println(l1)
	fmt.Println(l2)
	l3 := Or(l1, l2)
	fmt.Println(l3)
	l4 := Not(l3)
	fmt.Println(l4)
}
