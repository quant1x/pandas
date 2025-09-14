package pandas

import "gitee.com/quant1x/num"

// ScalarAggregation 标量聚合接口
type ScalarAggregation interface {
	// Mean calculates the average value of a series
	Mean() num.DType
	// StdDev calculates the standard deviation of a series
	StdDev() num.DType
	// Max 找出最大值
	Max() any
	// ArgMax Returns the indices of the maximum values along an axis
	ArgMax() int
	// Min 找出最小值
	Min() any
	// ArgMin Returns the indices of the minimum values along an axis
	ArgMin() int
	// Diff 元素的第一个离散差
	Diff(param any) (s Series)
	// Std 计算标准差
	Std() num.DType
	// Sum 计算累和
	Sum() num.DType
	// Add 加
	Add(x any) Series
	// Sub 减
	Sub(x any) Series
	// Mul 乘
	Mul(x any) Series
	// Div 除
	Div(x any) Series
	// Eq 等于
	Eq(x any) Series
	// Neq 不等于
	Neq(x any) Series
	// Gt 大于
	Gt(x any) Series
	// Gte 大于等于
	Gte(x any) Series
	// Lt 小于
	Lt(x any) Series
	// Lte 小于等于
	Lte(x any) Series
	// And 与
	And(x any) Series
	// Or 或
	Or(x any) Series
	// Not 非
	Not() Series
}
