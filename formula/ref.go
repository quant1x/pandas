package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// REF 引用前N的序列
func REF(S pandas.Series, N any) pandas.Series {
	return S.Ref(N)
}

// Deprecated: 推荐 REF [wangfeng on 2024/2/19 12:50]
func REF2[T num.BaseType](S []T, N any) []T {
	return num.Shift[T](S, N)
}
