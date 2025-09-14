package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
)

// MA 计算移动均线
//
//	求序列的N日简单移动平均值, 返回序列
func MA(S pandas.Series, N any) pandas.Series {
	return v2MA(S, N)
}

func v1MA(S pandas.Series, N any) pandas.Series {
	return S.Rolling(N).Mean()
}

func v2MA(S pandas.Series, N any) pandas.Series {
	x := S.DTypes()
	d := num.RollingV1(x, N, func(N num.DType, values ...float64) float64 {
		return num.Mean2(values)
	})
	return pandas.SliceToSeries(d)
}
