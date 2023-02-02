package pandas

import (
	"fmt"
	"testing"
)

func Test_point_to_1float32(t *testing.T) {
	var p1 *int8
	f1 := AnyToFloat32(p1)
	fmt.Printf("*int8 to float32=%f\n", f1)

	var v1 int8 = 1
	p1 = &v1
	f1 = AnyToFloat32(p1)
	fmt.Printf("*int8 to float32=%f\n", f1)
}

func Test_point_to_float32(t *testing.T) {
	type args struct {
		v any
	}

	// 指针声明
	var pInt8 *int8
	var pUint8 *uint8
	var pInt16 *int16
	var pUint16 *uint16
	var pInt32 *int32
	var pUint32 *uint32
	var pInt64 *int64
	var pUint64 *uint64
	var pInt *int
	var pUint *uint
	var pFloat32 *float32
	var pFloat64 *float64
	var pBoolTrue *bool
	var pBoolFalse *bool
	var pStr1 *string
	var pStr2 *string

	vInt8 := int8(1)
	pInt8 = &vInt8
	vUint8 := uint8(1)
	pUint8 = &vUint8

	vInt16 := int16(1)
	pInt16 = &vInt16
	vUint16 := uint16(1)
	pUint16 = &vUint16

	vInt32 := int32(1)
	pInt32 = &vInt32
	vUint32 := uint32(1)
	pUint32 = &vUint32

	vInt64 := int64(1)
	pInt64 = &vInt64
	vUint64 := uint64(1)
	pUint64 = &vUint64

	vInt := int(1)
	pInt = &vInt
	vUint := uint(1)
	pUint = &vUint

	vFloat32 := float32(1)
	pFloat32 = &vFloat32
	vFloat64 := float64(1)
	pFloat64 = &vFloat64

	vBoolTrue := true
	pBoolTrue = &vBoolTrue
	vBoolFalse := false
	pBoolFalse = &vBoolFalse

	vStr1 := "abc"
	pStr1 = &vStr1
	vStr2 := "1.23"
	pStr2 = &vStr2

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "T01: int8",
			args: args{
				pInt8,
			},
			want: float32(1),
		},
		{
			name: "T02: uint8",
			args: args{
				pUint8,
			},
			want: float32(1),
		},
		{
			name: "T03: int16",
			args: args{
				pInt16,
			},
			want: float32(1),
		},
		{
			name: "T04: uint16",
			args: args{
				pUint16,
			},
			want: float32(1),
		},
		{
			name: "T05: int32",
			args: args{
				pInt32,
			},
			want: float32(1),
		},
		{
			name: "T06: uint32",
			args: args{
				pUint32,
			},
			want: float32(1),
		},
		{
			name: "T07: int64",
			args: args{
				pInt64,
			},
			want: float32(1),
		},
		{
			name: "T08: uint64",
			args: args{
				pUint64,
			},
			want: float32(1),
		},
		{
			name: "T09: int",
			args: args{
				pInt,
			},
			want: float32(1),
		},
		{
			name: "T10: uint",
			args: args{
				pUint,
			},
			want: float32(1),
		},
		{
			name: "T11: float32",
			args: args{
				pFloat32,
			},
			want: float32(1),
		},
		{
			name: "T12: float64",
			args: args{
				pFloat64,
			},
			want: float32(1),
		},
		{
			name: "T13: true",
			args: args{
				pBoolTrue,
			},
			want: float32(1),
		},
		{
			name: "T14: false",
			args: args{
				pBoolFalse,
			},
			want: float32(0),
		},
		{
			name: "T15: str1",
			args: args{
				pStr1,
			},
			want: Nil2Float32,
		},
		{
			name: "T16: str2",
			args: args{
				pStr2,
			},
			want: float32(1.23),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got := point_to_float32(tt.args.v)
			//if got != tt.want {
			//	if !IsNaN(float64(tt.want)) {
			//		t.Errorf("point_to_float32() = %v, want %v", got, tt.want)
			//	} else if !IsNaN(float64(got)) {
			//		t.Errorf("point_to_float32() = %v, want %v", got, tt.want)
			//	}
			//}
			if got := AnyToFloat32(tt.args.v); got != tt.want && !(Float64IsNaN(float64(tt.want)) && Float64IsNaN(float64(got))) {
				t.Errorf("AnyToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_value_to_float32(t *testing.T) {
	type args struct {
		v any
	}

	vInt8 := int8(1)
	vUint8 := uint8(1)

	vInt16 := int16(1)
	vUint16 := uint16(1)

	vInt32 := int32(1)
	vUint32 := uint32(1)

	vInt64 := int64(1)
	vUint64 := uint64(1)

	vInt := int(1)
	vUint := uint(1)

	vFloat32 := float32(1)
	vFloat64 := float64(1)

	//vBoolTrue := true
	//vBoolFalse := false

	vStr1 := "abc"
	vStr2 := "1.23"

	// 组装测试用例
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "T01: int8",
			args: args{
				vInt8,
			},
			want: float32(1),
		},
		{
			name: "T02: uint8",
			args: args{
				vUint8,
			},
			want: float32(1),
		},
		{
			name: "T03: int16",
			args: args{
				vInt16,
			},
			want: float32(1),
		},
		{
			name: "T04: uint16",
			args: args{
				vUint16,
			},
			want: float32(1),
		},
		{
			name: "T05: int32",
			args: args{
				vInt32,
			},
			want: float32(1),
		},
		{
			name: "T06: uint32",
			args: args{
				vUint32,
			},
			want: float32(1),
		},
		{
			name: "T07: int64",
			args: args{
				vInt64,
			},
			want: float32(1),
		},
		{
			name: "T08: uint64",
			args: args{
				vUint64,
			},
			want: float32(1),
		},
		{
			name: "T09: int",
			args: args{
				vInt,
			},
			want: float32(1),
		},
		{
			name: "T10: uint",
			args: args{
				vUint,
			},
			want: float32(1),
		},
		{
			name: "T11: float32",
			args: args{
				vFloat32,
			},
			want: float32(1),
		},
		{
			name: "T12: float64",
			args: args{
				vFloat64,
			},
			want: float32(1),
		},
		{
			name: "T13: true",
			args: args{
				true,
			},
			want: float32(1),
		},
		{
			name: "T14: false",
			args: args{
				false,
			},
			want: float32(0),
		},
		{
			name: "T15: str1",
			args: args{
				vStr1,
			},
			want: Nil2Float32,
		},
		{
			name: "T16: str2",
			args: args{
				vStr2,
			},
			want: float32(1.23),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got := value_to_float32(tt.args.v)
			//if got != tt.want {
			//	if !IsNaN(float64(tt.want)) {
			//		t.Errorf("value_to_float32() = %v, want %v", got, tt.want)
			//	} else if !IsNaN(float64(got)) {
			//		t.Errorf("value_to_float32() = %v, want %v", got, tt.want)
			//	}
			//}
			if got := AnyToFloat32(tt.args.v); !(got == tt.want || (Float64IsNaN(float64(tt.want)) && Float64IsNaN(float64(got)))) {
				t.Errorf("AnyToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}
