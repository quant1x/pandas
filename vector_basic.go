package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
)

func (this Vector[T]) Name() string {
	return defaultSeriesName()
}

func (this Vector[T]) Rename(name string) {
	_ = name
}

func (this Vector[T]) Type() Type {
	return num.CheckoutRawType(this)
}

func (this Vector[T]) Values() any {
	return []T(this)
}

func (this Vector[T]) Reverse() Series {
	return api.Reverse(this)
}
