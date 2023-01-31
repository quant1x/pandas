package pandas

import (
	"math"
	"testing"
)

func TestParseFloat2Nan(t *testing.T) {
	type args struct {
		s string
		v any
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "NaN-str",
			args: args{
				s: "a",
				v: "a",
			},
			want: math.NaN(),
		},
		{
			name: "NaN-NaN",
			args: args{
				s: StringNaN,
				v: StringNaN,
			},
			want: math.NaN(),
		},
		{
			name: "NaN-nan",
			args: args{
				s: "nan",
				v: "nan",
			},
			want: math.NaN(),
		},
		{
			name: "NaN-Nan",
			args: args{
				s: "Nan",
				v: "Nan",
			},
			want: math.NaN(),
		},
		{
			name: "NaN-NAn",
			args: args{
				s: "NAn",
				v: "NAn",
			},
			want: math.NaN(),
		},
		{
			name: "NaN-NAN",
			args: args{
				s: "NAN",
				v: "NAN",
			},
			want: math.NaN(),
		},
		{
			name: "NaN-nAn",
			args: args{
				s: "nAn",
				v: "nAn",
			},
			want: math.NaN(),
		},
		{
			name: "NaN-nAN",
			args: args{
				s: "nAN",
				v: "nAN",
			},
			want: math.NaN(),
		},
		{
			name: "NaN-naN",
			args: args{
				s: "naN",
				v: "naN",
			},
			want: math.NaN(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := ParseFloat(tt.args.s, tt.args.v); got != tt.want {
			if got := ParseFloat(tt.args.s, tt.args.v); !math.IsNaN(got) {
				t.Errorf("ParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
