package pandas

import (
	"fmt"
	"testing"

	"gitee.com/quant1x/num"
)

func TestNDArray_Gt(t *testing.T) {
	d1 := []float32{11, 12, 13, 14, 15}
	d2 := []float64{1, 2, 3, 34, 5}

	s1 := vector[float32](d1)
	s2 := vector[float64](d2)

	fmt.Println("----------<Gt>----------")
	fmt.Println(s1.Gt(s2))
	fmt.Println(s1.Gt(13))
	fmt.Println("----------<Gte>----------")
	fmt.Println(num.Gte(d1, d2))
	fmt.Println(num.Gte(d1, 13))
	fmt.Println("----------<Lt>----------")
	fmt.Println(num.Lt(d1, d2))
	fmt.Println(num.Lt(d1, 13))
	fmt.Println("----------<Lte>----------")
	fmt.Println(num.Lte(d1, d2))
	fmt.Println(num.And(d1, 13))
	fmt.Println(num.Or(d1, 13))
}
