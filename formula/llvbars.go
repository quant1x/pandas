package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// LLVBARS 求上一低点到当前的周期数.
//
//	用法:
//	LLVBARS(X,N):求N周期内X最低值到当前周期数,N=0表示从第一个有效值开始统计
//	例如:
//	LLVBARS(HIGH,20)求得20日最低点到当前的周期数
func LLVBARS(S pandas.Series, N any) pandas.Series {
	x := S.Rolling(N).Apply(func(X pandas.Series, W num.DType) num.DType {
		return num.Any2DType(X.Reverse().ArgMin())
	})
	return x
}
