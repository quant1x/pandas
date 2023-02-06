package formula

import "gitee.com/quant1x/pandas"

type InputType interface {
	~int | ~[]float32 | ~[]float64
}

func HHV(S pandas.Series, N any) pandas.Series {
	return S.Rolling2(N).Max()
}
