package formula

import (
	"fmt"
	"slices"
	"testing"

	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

func TestMa_basic(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := pandas.Vector(x)
	r := MA(s, 5)
	fmt.Println(r)
}

func TestMA(t *testing.T) {
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
	B := df2.Col("b")

	// 2日均线
	r2 := MA(B, 2)
	fmt.Println(r2)
}

func BenchmarkMA_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkMA_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		r := MA(s, 5)
		_ = r
	}
}

func BenchmarkMA_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		r := v1MA(s, 5)
		_ = r
	}
}

func BenchmarkMA_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		r := v2MA(s, 5)
		_ = r
	}
}

func BenchmarkMA_v3Rolling_raw(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		d := num.RollingApply(s.DTypes(), 5, func(N num.DType, values ...float64) float64 {
			return num.Mean2(values)
		})
		s := pandas.SliceToSeries(d)
		_ = s
	}
}

func BenchmarkMA_v3Rolling(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		d := num.RollingApply(s.DTypes(), 5, func(N num.DType, values ...float64) float64 {
			return num.Mean2(values)
		})
		r := pandas.SliceToSeries(d)
		_ = r
	}
}
