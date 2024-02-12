package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// IF 序列布尔判断 return=A  if S==True  else  B
func IF(S stat.Series, A, B any) stat.Series {
	return IFF(S, A, B)
}

// IFF 序列布尔判断 return=A  if S==True  else  B
func IFF(S stat.Series, A, B any) stat.Series {
	length := S.Len()
	s := S.DTypes()

	a := stat.Align2Series(A, length).DTypes()
	b := stat.Align2Series(B, length).DTypes()
	ret := num.Where(s, a, b)
	return stat.ToSeries(ret...)
}

// IFN 序列布尔判断 return=A  if S==False  else  B
func IFN(S stat.Series, A, B any) stat.Series {
	return IFF(S, B, A)
}
