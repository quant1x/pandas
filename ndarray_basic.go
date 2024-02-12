package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
)

func (self NDArray[T]) Name() string {
	return "x"
}

func (self NDArray[T]) Rename(name string) {
	_ = name
}

func (self NDArray[T]) Type() Type {
	return num.CheckoutRawType(self)
}

func (self NDArray[T]) Values() any {
	return []T(self)
}

func (self NDArray[T]) Reverse() Series {
	return api.Reverse(self)
}
