package stat

import (
	"gitee.com/quant1x/pandas/algorithms/avx2"
	"math"
)

//func Where[T int64 | float64](condition any, params ...any) []T {
//	c1, c2, err := detectParam[float64](condition)
//	if err != nil {
//		return []T{}
//	}
//
//}

// Where 返回根据“条件”从“x”或“y”中选择的元素
// 这里先实现一个简单的, 留给于总重构
// params只支持两个默认值x和y, 如果condition为true返回x, 否则返回y
// condition和param都可能是基础数据类型,也可能是一个slice, 并且长度可能不一致
// 直接写成序列版本, 可能更简单
// func Where[T int64 | float64](condition []T, x, y []T) []T {
func Where1(condition []float64, x, y []float64) []float64 {
	// 第一步, 找出最长的
	clen := len(condition)
	xlen := len(x)
	ylen := len(y)
	// 第二步, 找出最大长度
	c := []float64{float64(clen), float64(xlen), float64(ylen)}
	maxLength := int(avx2.Max(c))

	// 对齐所有长度
	if clen < maxLength {
		condition = align(condition, math.NaN(), maxLength)
	}
	if xlen < maxLength {
		x = align(x, math.NaN(), maxLength)
	}
	if ylen < maxLength {
		y = align(y, math.NaN(), maxLength)
	}
	// 初始化返回值
	d := make([]float64, maxLength)
	for i := 0; i < maxLength; i++ {
		// NaN也被认为是真
		if condition[i] != 0 {
			d[i] = x[i]
		} else {
			d[i] = y[i]
		}
	}
	return d
}

func Where[T StatType](condition []T, x, y []T) []T {
	// 第一步, 找出最长的
	clen := len(condition)
	xlen := len(x)
	ylen := len(y)
	// 第二步, 找出最大长度
	c := []float64{float64(clen), float64(xlen), float64(ylen)}
	maxLength := int(avx2.Max(c))

	// 处理默认值
	defaultValue := typeDefault(T(0))
	// 对齐所有长度
	if clen < maxLength {
		condition = align(condition, defaultValue, maxLength)
	}
	if xlen < maxLength {
		x = align(x, defaultValue, maxLength)
	}
	if ylen < maxLength {
		y = align(y, defaultValue, maxLength)
	}
	// 初始化返回值
	d := make([]T, maxLength)
	for i := 0; i < maxLength; i++ {
		// NaN也被认为是真
		if condition[i] != 0 {
			d[i] = x[i]
		} else {
			d[i] = y[i]
		}
	}
	return d
}
