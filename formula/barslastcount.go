package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// BARSLASTCOUNT 统计连续满足S条件的周期数
func BARSLASTCOUNT(S stat.Series) stat.Series {
	s := S.DTypes()
	slen := len(s)
	rt := stat.Repeat[stat.Int](0, slen+1)
	for i := 0; i < slen; i++ {
		if s[i] != 0 {
			rt[i+1] = rt[i] + 1
		} else {
			rt[i+1] = rt[i+1]
		}
	}
	return stat.NewSeries[stat.Int](rt[1:]...)
}
