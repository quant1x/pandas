package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// BARSLAST 上一次条件成立到当前的周期, BARSLAST(C/REF(C,1)>=1.1) 上一次涨停到今天的天数
//
//	为了测试SMA,BARSLAST必须要先实现, 给SMA提供序列换参数, 以便验证, python那边还没实现
func BARSLAST(S stat.Series) stat.Series {
	ns := BARSLAST2(S)
	return stat.ToSeries(ns...)
}

func BARSLAST2(S stat.Series) []num.DType {
	fs := S.DTypes()
	as := num.Repeat[num.DType](1, S.Len())
	bs := num.Repeat[num.DType](0, S.Len())
	x := num.Where(fs, as, bs)
	M := []num.DType{0}
	M = append(M, x...)
	for i := 1; i < len(M); i++ {
		if int(M[i]) != 0 {
			M[i] = 0
		} else {
			M[i] = M[i-1] + 1
		}
	}
	return M[1:]
}
