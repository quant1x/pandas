package stat

import "testing"

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
			name: "empty",
			args: args{s: ""},
			want: true,
		},
		{
			name: "empty",
			args: args{s: "\t"},
			want: true,
		},
		{
			name: "NaN",
			args: args{s: StringNaN},
			want: false,
		},
		{
			name: "abc",
			args: args{s: StringNaN},
			want: false,
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
