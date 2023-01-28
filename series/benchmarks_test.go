package series_test

import (
	"math/rand"
	"strconv"
	"testing"

	"gitee.com/quant1x/pandas/series"
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
		seriesType series.Type
	}{
		{
			"[]bool(100000)_Int",
			seriesGenerateBools(100000),
			series.Int,
		},
		{
			"[]bool(100000)_String",
			seriesGenerateBools(100000),
			series.String,
		},
		{
			"[]bool(100000)_Bool",
			seriesGenerateBools(100000),
			series.Bool,
		},
		{
			"[]bool(100000)_Float",
			seriesGenerateBools(100000),
			series.Float,
		},
		{
			"[]string(100000)_Int",
			seriesGenerateStrings(100000),
			series.Int,
		},
		{
			"[]string(100000)_String",
			seriesGenerateStrings(100000),
			series.String,
		},
		{
			"[]string(100000)_Bool",
			seriesGenerateStrings(100000),
			series.Bool,
		},
		{
			"[]string(100000)_Float",
			seriesGenerateStrings(100000),
			series.Float,
		},
		{
			"[]float64(100000)_Int",
			seriesGenerateFloats(100000),
			series.Int,
		},
		{
			"[]float64(100000)_String",
			seriesGenerateFloats(100000),
			series.String,
		},
		{
			"[]float64(100000)_Bool",
			seriesGenerateFloats(100000),
			series.Bool,
		},
		{
			"[]float64(100000)_Float",
			seriesGenerateFloats(100000),
			series.Float,
		},
		{
			"[]int(100000)_Int",
			seriesGenerateInts(100000),
			series.Int,
		},
		{
			"[]int(100000)_String",
			seriesGenerateInts(100000),
			series.String,
		},
		{
			"[]int(100000)_Bool",
			seriesGenerateInts(100000),
			series.Bool,
		},
		{
			"[]int(100000)_Float",
			seriesGenerateInts(100000),
			series.Float,
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				series.NewSeries(test.data, test.seriesType, test.name)
			}
		})
	}
}

func BenchmarkSeries_Copy(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name   string
		series series.Series
	}{
		{
			"[]int(100000)_Int",
			series.Ints(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_String",
			series.Strings(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			series.Bools(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			series.Floats(seriesGenerateInts(100000)),
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
		series  series.Series
	}{
		{
			"[]int(100000)_Int",
			seriesGenerateIntsN(10000, 2),
			series.Ints(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_String",
			seriesGenerateIntsN(10000, 2),
			series.Strings(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			seriesGenerateIntsN(10000, 2),
			series.Bools(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			seriesGenerateIntsN(10000, 2),
			series.Floats(seriesGenerateInts(100000)),
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
		newValues series.Series
		series    series.Series
	}{
		{
			"[]int(100000)_Int",
			seriesGenerateIntsN(10000, 2),
			series.Ints(seriesGenerateIntsN(10000, 2)),
			series.Ints(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_String",
			seriesGenerateIntsN(10000, 2),
			series.Strings(seriesGenerateIntsN(10000, 2)),
			series.Strings(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			seriesGenerateIntsN(10000, 2),
			series.Bools(seriesGenerateIntsN(10000, 2)),
			series.Bools(seriesGenerateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			seriesGenerateIntsN(10000, 2),
			series.Floats(seriesGenerateIntsN(10000, 2)),
			series.Floats(seriesGenerateInts(100000)),
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
