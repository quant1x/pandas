package stat

import (
	"fmt"
	"testing"
)

func TestDiff(t *testing.T) {
	d1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(d1)
	fmt.Println("------------------------------------------------------------")
	N := 1
	fmt.Println("固定的参数, N =", N)
	r1 := Diff(d1, N)
	fmt.Println("序列化结果:", r1)
	fmt.Println("------------------------------------------------------------")
	s1 := []float64{1, 2, 3, 4, 3, 3, 2, 1, Nil2Float64, Nil2Float64, Nil2Float64, Nil2Float64}
	fmt.Printf("序列化参数: %+v\n", s1)
	r2 := Diff(d1, s1)
	fmt.Println("序列化结果:", r2)
}
