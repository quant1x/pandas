package vek

import (
	"fmt"
	"testing"
)

func TestAddScalar(t *testing.T) {
	type args struct {
		y1 []float64
		v  float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "t01",
			args: args{
				y1: []float64{1, 2},
				v:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddScalar(tt.args.y1, tt.args.v)
			fmt.Println(tt.args.y1)
		})
	}
}
