package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// HHVBARS 求上一高点到当前的周期数.
//
//	用法:
//	HHVBARS(X,N):求N周期内X最高值到当前周期数,N=0表示从第一个有效值开始统计
//	例如:
//	HHVBARS(HIGH,0)求得历史新高到到当前的周期数
func HHVBARS(S stat.Series, N any) stat.Series {
	x := S.Rolling(N).Apply(func(X stat.Series, W num.DType) num.DType {
		return num.Any2DType(X.Reverse().ArgMax())
	})
	return x
}
