package utils

import (
	"math"
	"reflect"
	"unsafe"
)

func WantFloat(got, want float64) bool {
	return got != want && !(math.IsNaN(want) && math.IsNaN(got))
}

func SliceWantFloat(got, want []float64) bool {
	b := 0
	for i := 0; i < len(got); i++ {
		b1 := got[i] == want[i] || (math.IsNaN(want[i]) && math.IsNaN(got[i]))
		if b1 {
			b += 1
		}
	}
	return b == len(got)
}

// ChanIsClosed 判断channel是否关闭
func ChanIsClosed(ch any) bool {
	if reflect.TypeOf(ch).Kind() != reflect.Chan {

		panic("only channels!")

	}
	cptr := *(*uintptr)(unsafe.Pointer(
		unsafe.Pointer(uintptr(unsafe.Pointer(&ch)) + unsafe.Sizeof(uint(0))),
	))
	// this function will return true if chan.closed > 0
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32
	// **
	cptr += unsafe.Sizeof(uint(0)) * 2
	cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
	cptr += unsafe.Sizeof(uint16(0))
	return *(*uint32)(unsafe.Pointer(cptr)) > 0
}
