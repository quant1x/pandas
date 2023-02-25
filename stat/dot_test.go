package stat

import (
	"fmt"
	"testing"
)

func TestDot2D(t *testing.T) {
	// https://blog.csdn.net/llittleSun/article/details/115045660
	//
	//a := [][]int{{1, 2}, {3, 7}}
	//b := [][]int{{4, 3}, {5, 0}}
	//
	//c := Dot2D(a, b)
	//fmt.Println(c)

	A := [][]int{{1, 4, 9}, {1, 2, 3}, {1, 1, 1}}
	fmt.Println("A =", A)
	B := [][]int{{1, 1, 1}, {4, 2, 1}, {9, 3, 1}}
	fmt.Println("B =", B)
	C := Dot2D(A, B)
	fmt.Println("C =", C)
}

func TestDot(t *testing.T) {
	A := [][]int{{1, 4, 9}, {1, 2, 3}, {1, 1, 1}}
	A = [][]int{{1, 4, 9, 16, 25}, {1, 2, 3, 4, 5}, {1, 1, 1, 1, 1}}
	B := Transpose2D(A)
	fmt.Println("A =", A)
	fmt.Println("B =", B)
	C := Dot(A, B)
	fmt.Println("C =", C)
}
