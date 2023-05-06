package internal

import (
	"math"
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
