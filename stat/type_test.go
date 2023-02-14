package stat

import (
	"fmt"
	"testing"
)

func Test_typeDefault(t *testing.T) {
	fmt.Println(typeDefault[float32]())
	fmt.Println(typeDefault[float64]())

	fmt.Println(typeDefault[bool]())
	fmt.Println(typeDefault[string]())
	fmt.Println(typeDefault[uintptr]())
	fmt.Println(typeDefault[int8]())
	fmt.Println(typeDefault[uint8]())
	fmt.Println(typeDefault[int16]())
	fmt.Println(typeDefault[uint16]())
	fmt.Println(typeDefault[int32]())
	fmt.Println(typeDefault[uint32]())
	fmt.Println(typeDefault[int64]())
	fmt.Println(typeDefault[uint64]())
	fmt.Println(typeDefault[int]())
	fmt.Println(typeDefault[uint]())
}

func Test_Number(t *testing.T) {

}

func Test_anyToGeneric(t *testing.T) {
	fmt.Println(anyToGeneric[int](true))
	fmt.Println(anyToGeneric[int]("true"))
	fmt.Println(anyToGeneric[int]("false"))
	fmt.Println(anyToGeneric[int]("aa"))
	fmt.Println(anyToGeneric[int]("tt"))
	fmt.Println(anyToGeneric[int](3.00))
}
