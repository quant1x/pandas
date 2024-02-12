package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// BARSSINCEN N周期内第一次S条件成立到现在的周期数,N为常量
func BARSSINCEN(S pandas.Series, N any) pandas.Series {
	ret := S.Rolling(N).Apply(func(X pandas.Series, M num.DType) num.DType {
		x := X.DTypes()
		n := int(M)
		argMax := num.ArgMax(x)
		r := 0
		if argMax != 0 || x[0] != 0 {
			r = n - 1 - argMax
		} else {
			r = 0
		}
		return num.DType(r)
	})
	r1 := ret.FillNa(0, true)
	return r1
}
