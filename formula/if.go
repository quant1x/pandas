package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// IF 序列布尔判断 return=A  if S==True  else  B
func IF(S pandas.Series, A, B any) pandas.Series {
	return IFF(S, A, B)
}

// IFF 序列布尔判断 return=A  if S==True  else  B
func IFF(S pandas.Series, A, B any) pandas.Series {
	length := S.Len()
	s := S.DTypes()

	a := pandas.Align2Series(A, length).DTypes()
	b := pandas.Align2Series(B, length).DTypes()
	ret := num.Where(s, a, b)
	return pandas.SliceToSeries(ret)
}

// IFN 序列布尔判断 return=A  if S==False  else  B
func IFN(S pandas.Series, A, B any) pandas.Series {
	return IFF(S, B, A)
}
