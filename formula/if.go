package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// IF 序列布尔判断 return=A  if S==True  else  B
func IF(S, A, B stat.Series) stat.Series {
	return IFF(S, A, B)
}

// IFF 序列布尔判断 return=A  if S==True  else  B
func IFF(S, A, B stat.Series) stat.Series {
	s := S.DTypes()
	a := A.DTypes()
	b := B.DTypes()
	ret := stat.Where(s, a, b)
	return stat.NewSeries[stat.DType](ret...)
}

// IFN 序列布尔判断 return=A  if S==False  else  B
func IFN(S, A, B stat.Series) stat.Series {
	return IFF(S, B, A)
}
