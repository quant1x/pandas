package pandas

import (
	"github.com/quant1x/num"
	"github.com/quant1x/x/api"
)

func (this vector[T]) Name() string {
	return defaultSeriesName()
}

func (this vector[T]) Rename(name string) {
	_ = name
}

func (this vector[T]) Type() Type {
	return num.CheckoutRawType(this)
}

func (this vector[T]) Values() any {
	return []T(this)
}

func (this vector[T]) Reverse() Series {
	return api.Reverse(this)
}
