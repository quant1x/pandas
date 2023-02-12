package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// COUNT 统计S为真的天数
func COUNT(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Count()
}
func COUNT2(S []bool, N int) []int {
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

// COUNT_v1 一般性比较
func COUNT_v1(S pandas.Series, N any) []stat.Int {
	//values := S.DTypes()
	return S.Rolling(N).Apply(func(X pandas.Series, W stat.DType) stat.DType {
		x := X.DTypes()
		n := 0
		for _, v := range x {
			if v != 0 {
				n++
			}
		}
		return stat.DType(n)
	}).AsInt()
}
