package stat

import (
	"gitee.com/quant1x/vek"
)

type DType = float64
type Int = int32

// DTypeIsNaN 判断DType是否NaN
func DTypeIsNaN(d DType) bool {
	return Float64IsNaN(d)
}

// Slice2DType 切片转DType
func Slice2DType(v any) []DType {
	return SliceToFloat64(v)
}

// Any2DType any转DType
func Any2DType(v any) DType {
	return AnyToFloat64(v)
}

// DType切片转int32切片
func DType2Int(d []DType) []Int {
	return vek.ToInt32(d)
}
