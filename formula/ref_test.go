package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/num"
	"github.com/quant1x/num/labs"
	"github.com/quant1x/pandas"
)

func TestREF_basic(t *testing.T) {
	C := pandas.Vector([]float64{1, 2, 3, 4, 5})
	fmt.Println(C)
	P := 1
	fmt.Println("常量 =", P)
	s1 := REF(C, P)
	fmt.Println("\t=>", s1)
	N := []int{1, 2, 3, 2, 1}
	fmt.Println("可变参数 =", N)
	s2 := REF(C, N)
	fmt.Println("\t=>", s2)
}

func TestREF(t *testing.T) {
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
			name: "ref-const",
			args: args{
				S: pandas.Vector([]float64{1, 2, 3, 4, 5}),
				N: 1,
			},
			want: pandas.Vector([]float64{num.Float64NaN(), 1, 2, 3, 4}),
		},
		{
			name: "ref-vector-1",
			args: args{
				S: pandas.Vector([]float64{1, 2, 3, 4, 5}),
				N: []int{1, 1, 1, 1, 1},
			},
			want: pandas.Vector([]float64{num.Float64NaN(), 1, 2, 3, 4}),
		},
		{
			name: "ref-vector-x",
			args: args{
				S: pandas.Vector([]float64{1, 2, 3, 4, 5}),
				N: []int{1, 2, 3, 1, 0},
			},
			want: pandas.Vector([]float64{num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := REF(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("REF() = %v, want %v", got, tt.want)
			}
		})
	}
}
