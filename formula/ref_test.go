package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestREF_basic(t *testing.T) {
	C := pandas.Vector([]float64{1, 2, 3, 4, 5})
	fmt.Println(C)
	P := 1
	fmt.Println("常量 =", P)
	s1 := REF(C, P)
	fmt.Println("\t=>", s1)
	N := []int{1, 2, 3, 2, 1}
	fmt.Println("可变参数 =", N)
	s2 := REF(C, N)
	fmt.Println("\t=>", s2)
}

func TestREF(t *testing.T) {
	type testStruct struct {
		A string
		B int
		C bool
		D float32
	}
	data := []testStruct{
		{"a", 1, true, 0.0},
		{"b", 2, false, 0.5},
	}
	df1 := pandas.LoadStructs(data)
	fmt.Println(df1)
	// 修改列名
	_ = df1.SetNames("a", "b", "c", "d")
	// 增加1列
	s_e := pandas.NewSeriesWithoutType("", "a0", "a1", "a2", "a3")
	df2 := df1.Join(s_e)
	fmt.Println(df2)
	A := df2.Col("a")
	B := df2.Col("b")
	C := df2.Col("c")
	D := df2.Col("d")

	// 0日前的D值
	r0 := REF(D, 0)
	fmt.Println(r0)
	// 1日前的D值
	r1 := REF(D, 1)
	fmt.Println(r1)

	_ = A
	_ = B
	_ = C
	_ = D
}
