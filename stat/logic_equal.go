package stat

import (
	"gitee.com/quant1x/gox/num"
	"gitee.com/quant1x/gox/num/num32"
)

// Equal 比较相等
func Equal[T BaseType](x, y []T) []bool {
	return binaryOperations2[T, bool](x, y, num32.Eq, num.Eq, __eq_go[T])
}

func __eq_go[T BaseType](x, y []T) []bool {
	length := len(x)
	d := make([]bool, length)
	for i := 0; i < length; i++ {
		if x[i] == y[i] {
			d[i] = true
		}
	}
	return d
}
