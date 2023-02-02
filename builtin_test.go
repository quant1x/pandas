package pandas

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test: 字符串左右两边空格",
			args: args{
				s: " a ",
			},
			want: false,
		},
		{
			name: "test: 单空格",
			args: args{
				s: " ",
			},
			want: true,
		},
		{
			name: "test: 双空格",
			args: args{
				s: "  ",
			},
			want: true,
		},
		{
			name: "test: 多空格",
			args: args{
				s: "     ",
			},
			want: true,
		},
		{
			name: "test: tab",
			args: args{
				s: "\t",
			},
			want: true,
		},
		{
			name: "test: 空格和tab",
			args: args{
				s: " \t",
			},
			want: true,
		},
		{
			name: "test: tab和空格",
			args: args{
				s: " \t",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint(t *testing.T) {
	var p1 *int
	fmt.Printf("*int = nil, result=%v\n", isPoint(p1))
	a := 1
	p1 = &a
	fmt.Printf("*int = nil, result=%v\n", isPoint(p1))

	var p2 *BigFloat
	fmt.Printf("*BigFloat = nil, result=%v\n", isPoint(p2))

}
