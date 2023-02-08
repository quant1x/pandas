package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// BARSSINCEN N周期内第一次S条件成立到现在的周期数,N为常量
func BARSSINCEN(S pandas.Series, N any) []stat.Int {
	ret := S.Rolling(N).Apply(func(X pandas.Series, M stat.DType) stat.DType {
		x := X.DTypes()
		n := int(M)
		argMax := stat.ArgMax(x)
		r := 0
		if argMax != 0 || x[0] != 0 {
			r = n - 1 - argMax

		} else {
			r = 0
		}
		return stat.DType(r)
	})
	r1 := ret.FillNa(0, true)
	r2 := r1.AsInt()
	return r2
}
