package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// BARSLASTCOUNT 统计连续满足S条件的周期数
func BARSLASTCOUNT(S pandas.Series) pandas.Series {
	s := S.DTypes()
	slen := len(s)
	rt := num.Repeat[int](0, slen+1)
	for i := 0; i < slen; i++ {
		if s[i] != 0 {
			rt[i+1] = rt[i] + 1
		} else {
			rt[i+1] = rt[i+1]
		}
	}
	ns := rt[1:]
	return pandas.SliceToSeries(ns)
}
