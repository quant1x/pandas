package stat

import (
	"fmt"
	"testing"
)

func Test_integer2Bool(t *testing.T) {
	fmt.Println(integer2Bool(int(0)))
	fmt.Println(integer2Bool(int(1)))
	fmt.Println(integer2Bool(int64(0)))
	fmt.Println(integer2Bool(int64(1)))
	fmt.Println(integer2Bool(float32(0)))
	fmt.Println(integer2Bool(float32(1)))
	fmt.Println(integer2Bool(float64(0)))
	fmt.Println(integer2Bool(float64(1)))
}
