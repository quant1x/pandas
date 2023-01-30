package dframe

import (
	"math/rand"
	"strconv"
	"testing"
)

func seriesGenerateInts(n int) (data []int) {
	for i := 0; i < n; i++ {
		data = append(data, rand.Int())
	}
	return
}

func seriesGenerateFloats(n int) (data []float64) {
	for i := 0; i < n; i++ {
		data = append(data, rand.Float64())
	}
	return
}

func seriesGenerateStrings(n int) (data []string) {
	for i := 0; i < n; i++ {
		data = append(data, strconv.Itoa(rand.Int()))
	}
	return
}

func seriesGenerateBools(n int) (data []bool) {
	for i := 0; i < n; i++ {
		r := rand.Intn(2)
		b := false
		if r == 1 {
			b = true
		}
		data = append(data, b)
	}
	return
}

func seriesGenerateIntsN(n, k int) (data []int) {
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(k))
	}
	return
}

func BenchmarkSeries_New(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name       string
		data       interface{}
		seriesType Type
	}{
		{
			"[]bool(100000)_Int",
			seriesGenerateBools(100000),
			Int,
		},
		{
			"[]bool(100000)_String",
			seriesGenerateBools(100000),
			String,
		},
		{
			"[]bool(100000)_Bool",
			seriesGenerateBools(100000),
			Bool,
		},
		{
			"[]bool(100000)_Float",
			seriesGenerateBools(100000),
			Float,
		},
		{
			"[]string(100000)_Int",
			seriesGenerateStrings(100000),
			Int,
		},
		{
			"[]string(100000)_String",
			seriesGenerateStrings(100000),
			String,
		},
		{
			"[]string(100000)_Bool",
			seriesGenerateStrings(100000),
			Bool,
		},
		{
			"[]string(100000)_Float",
			seriesGenerateStrings(100000),
			Float,
		},
		{
			"[]float64(100000)_Int",
			seriesGenerateFloats(100000),
			Int,
		},
		{
			"[]float64(100000)_String",
			seriesGenerateFloats(100000),
			String,
		},
		{
			"[]float64(100000)_Bool",
			seriesGenerateFloats(100000),
			Bool,
		},
		{
			"[]float64(100000)_Float",
			seriesGenerateFloats(100000),
			Float,
		},
		{
			"[]int(100000)_Int",
			seriesGenerateInts(100000),
			Int,
		},
		{
			"[]int(100000)_String",
			seriesGenerateInts(100000),
			String,
		},
		{
			"[]int(100000)_Bool",
			seriesGenerateInts(100000),
			Bool,
		},
		{
			"[]int(100000)_Float",
			seriesGenerateInts(100000),
			Float,
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NewSeries(test.data, test.seriesType, test.name)
			}
		})
	}
}

func BenchmarkSeries_Copy(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name   string
		series Series
	}{
		{
			"[]int(100000)_Int",
			Ints(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_String",
			Strings(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			Bools(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			Floats(seriesGenerateInts(100000)),
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.series.Copy()
			}
		})
	}
}

func BenchmarkSeries_Subset(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name    string
		indexes interface{}
		series  Series
	}{
		{
			"[]int(100000)_Int",
			seriesGenerateIntsN(10000, 2),
			Ints(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_String",
			seriesGenerateIntsN(10000, 2),
			Strings(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			seriesGenerateIntsN(10000, 2),
			Bools(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			seriesGenerateIntsN(10000, 2),
			Floats(seriesGenerateInts(100000)),
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.series.Subset(test.indexes)
			}
		})
	}
}

func BenchmarkSeries_Set(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name      string
		indexes   interface{}
		newValues Series
		series    Series
	}{
		{
			"[]int(100000)_Int",
			seriesGenerateIntsN(10000, 2),
			Ints(seriesGenerateIntsN(10000, 2)),
			Ints(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_String",
			seriesGenerateIntsN(10000, 2),
			Strings(seriesGenerateIntsN(10000, 2)),
			Strings(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			seriesGenerateIntsN(10000, 2),
			Bools(seriesGenerateIntsN(10000, 2)),
			Bools(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			seriesGenerateIntsN(10000, 2),
			Floats(seriesGenerateIntsN(10000, 2)),
			Floats(seriesGenerateInts(100000)),
		},
	}
	for _, test := range table {
		s := test.series.Copy()
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Set(test.indexes, test.newValues)
			}
		})
	}
}
