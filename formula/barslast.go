package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// BARSLAST 上一次条件成立到当前的周期, BARSLAST(C/REF(C,1)>=1.1) 上一次涨停到今天的天数
//
//	为了测试SMA,BARSLAST必须要先实现, 给SMA提供序列换参数, 以便验证, python那边还没实现
func BARSLAST(S pandas.Series) pandas.Series {
	ns := BARSLAST2(S)
	return pandas.SliceToSeries(ns)
}

func BARSLAST2(S pandas.Series) []num.DType {
	fs := S.DTypes()
	as := num.Repeat[num.DType](1, S.Len())
	bs := num.Repeat[num.DType](0, S.Len())
	x := num.Where(fs, as, bs)
	M := []num.DType{0}
	M = append(M, x...)
	last_true_pos := -1
	for i := 1; i < len(M); i++ {
		current := int(M[i]) != 0
		if current {
			M[i] = 0
			last_true_pos = i
		} else {
			if last_true_pos >= 0 {
				M[i] = num.DType(i - last_true_pos)
			} else {
				M[i] = -1
			}
		}
	}
	return M[1:]
}
