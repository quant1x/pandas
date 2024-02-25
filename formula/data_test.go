package formula

import (
	"math/rand"
	"sync"
)

const (
	benchAlignLength  = 5000
	benchAlignInitNum = 5000
)

var (
	testDataOnce    sync.Once
	testDataFloat32 []float32
	testDataFloat64 []float64
)

func initTestData() {
	testDataFloat32 = make([]float32, benchAlignInitNum)
	testDataFloat64 = make([]float64, benchAlignInitNum)
	for i := 0; i < benchAlignInitNum; i++ {
		testDataFloat32[i] = rand.Float32()
		testDataFloat64[i] = rand.Float64()
	}
}
