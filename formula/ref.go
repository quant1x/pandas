package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// REF 引用前N的序列
func REF(S stat.Series, N any) stat.Series {
	return S.Ref(N)
}

func REF2[T num.BaseType](S []T, N any) []T {
	return num.Shift[T](S, N)
}
