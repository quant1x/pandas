package pandas

import (
	"math"
	"testing"
)

func TestAnyToString(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test: float nan",
			args: args{
				v: math.NaN(),
			},
			want: StringNaN,
		},
		{
			name: "test: float true",
			args: args{
				v: true,
			},
			want: True2String,
		},
		{
			name: "test: float false",
			args: args{
				v: true,
			},
			want: False2String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyToString(tt.args.v); got != tt.want {
				t.Errorf("AnyToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
