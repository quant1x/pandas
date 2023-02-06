package pandas

import (
	"fmt"
	"math"
	"sort"
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
			name: "test: true true",
			args: args{
				v: true,
			},
			want: True2String,
		},
		{
			name: "test: false false",
			args: args{
				v: false,
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

func TestStringListSort(t *testing.T) {
	// 准备一个内容被打乱顺序的字符串切片
	names := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sa := Arraystring(names)
	// 使用sort包进行排序
	sort.Sort(sa)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
	n1 := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(n1)
	fmt.Println(n1)
}
