package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// IF 序列布尔判断 return=A  if S==True  else  B
func IF(S, A, B stat.Series) stat.Series {
	return IFF(S, A, B)
}

// IFF 序列布尔判断 return=A  if S==True  else  B
func IFF(S, A, B stat.Series) stat.Series {
	s := S.Floats()
	a := A.Floats()
	b := B.Floats()
	ret := stat.Where(s, a, b)
	return pandas.NewSeries(stat.SERIES_TYPE_FLOAT32, "", ret)
}

// IFN 序列布尔判断 return=A  if S==False  else  B
func IFN(S, A, B stat.Series) stat.Series {
	return IFF(S, B, A)
}
