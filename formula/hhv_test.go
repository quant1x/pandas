package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/num/labs"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestHHV_basic(t *testing.T) {

}

func TestHHV(t *testing.T) {
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
			name: "hhv",
			args: args{
				S: pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "x", []float64{2, 1, 3, 4, 5, 2}),
				N: 3,
			},
			want: pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT64, "x", []float64{num.Float64NaN(), num.Float64NaN(), 3, 4, 5, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HHV(tt.args.S, tt.args.N); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("HHV() = %v, want %v", got, tt.want)
			}
		})
	}
}
