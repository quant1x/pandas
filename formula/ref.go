package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// REF 引用前N的序列
func REF(S stat.Series, N any) stat.Series {
	return S.Ref(N)
}

func REF2[T stat.BaseType](S []T, N any) []T {
	return stat.Shift[T](S, N)
}
