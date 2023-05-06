package stat

import (
	"gitee.com/quant1x/pandas/internal"
	"math"
	"testing"
)

func TestWhere(t *testing.T) {
	type args struct {
		condition []float64
		x         []float64
		y         []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "t01",
			args: args{
				condition: []float64{1, 1, 1},
				x:         []float64{0.1, 0.2, 0.3},
				y:         []float64{1.1, 1.2, 1.3},
			},
			want: []float64{0.1, 0.2, 0.3},
		},
		{
			name: "t02",
			args: args{
				condition: []float64{1, 0},
				x:         []float64{0.1, 0.2, 0.3},
				y:         []float64{1.1, 1.2, 1.3},
			},
			want: []float64{0.1, 1.2, 0.3},
		},
		{
			name: "t03",
			args: args{
				condition: []float64{1, 0},
				x:         []float64{0.1, 0.2},
				y:         []float64{1.1, math.NaN(), 1.3},
			},
			want: []float64{0.1, math.NaN(), math.NaN()},
		},
		{
			name: "t04",
			args: args{
				condition: []float64{1, 0, 1, 1},
				x:         []float64{0.1, 0.2},
				y:         []float64{1.1, math.NaN(), 1.3},
			},
			want: []float64{0.1, math.NaN(), math.NaN(), math.NaN()},
		},
		{
			name: "t05",
			args: args{
				condition: []float64{1, 0, 1},
				x:         []float64{0.1, 0.2, 0.3, 0.4},
				y:         []float64{1.1, math.NaN(), 1.3},
			},
			want: []float64{0.1, math.NaN(), 0.3, 0.4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Where(tt.args.condition, tt.args.x, tt.args.y); !internal.SliceWantFloat(got, tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}
