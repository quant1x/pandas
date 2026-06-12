package formula

import (
	"fmt"
	"slices"
	"testing"

	"github.com/quant1x/num"
	"github.com/quant1x/num/labs"
	"github.com/quant1x/pandas"
)

func TestLLV_basic(t *testing.T) {
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

	r2 := LLV(D, 2)
	fmt.Println(r2)

	r2 = LLV(A, 2)
	fmt.Println(r2)

	r2 = LLV(df2.Col("X0"), 2)
	fmt.Println(r2)

	_ = A
	_ = B
	_ = C
	_ = D
}

func TestLLV(t *testing.T) {
	type args struct {
		S pandas.Series
		N any
	}
	tests := []struct {
		name string
		args args
		want pandas.Series
	}{
		{
			name: "llv",
			args: args{
				S: pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "x", []float64{2, 1, 3, 4, 5, 2}),
				N: 3,
			},
			want: pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "x", []float64{num.Float64NaN(), num.Float64NaN(), 1, 1, 3, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LLV(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("LLV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLLV_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkLLV_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		LLV(s, 10)
	}
}

func BenchmarkLLV_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v1LLV(s, 10)
	}
}

func BenchmarkLLV_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := pandas.SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		v2LLV(s, 10)
	}
}
