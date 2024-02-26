package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// REF 引用前N的序列
func REF(S pandas.Series, N any) pandas.Series {
	return v1REF(S, N)
}

func v1REF(S pandas.Series, N any) pandas.Series {
	return S.Ref(N)
}

func v2REF[T num.BaseType](S []T, N any) []T {
	return num.Shift[T](S, N)
}
