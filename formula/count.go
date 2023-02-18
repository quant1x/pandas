package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// COUNT 统计S为真的天数
func COUNT(S []bool, N int) []int {
	xLen := len(S)
	x := stat.Rolling(S, N)
	ret := make([]int, xLen)
	for i := 0; i < len(x); i++ {
		n := 0
		for _, v := range x[i] {
			if stat.AnyToBool(v) {
				n++
			}
		}
		ret[i] = n
	}
	return ret
}
