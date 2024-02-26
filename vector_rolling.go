package pandas

import "gitee.com/quant1x/num"

func (this vector[T]) Rolling(param any) RollingAndExpandingMixin {
	window := num.Any2Window[num.DType](param)
	w := RollingAndExpandingMixin{
		Window: window,
		Series: this,
	}
	return w
}

func (this vector[T]) v1Rolling(param any) RollingAndExpandingMixin {
	window := num.Window[num.DType]{C: num.NaN()}
	switch v := param.(type) {
	case int:
		window.C = num.DType(v)
	case []num.DType:
		window.V = v
	case Series:
		vs := v.DTypes()
		window.V = vs
	default:
		panic(num.ErrInvalidWindow)
	}
	w := RollingAndExpandingMixin{
		Window: window,
		Series: this,
	}
	return w
}

func (this vector[T]) v2Rolling(param any) RollingAndExpandingMixin {
	var N []num.DType
	switch v := param.(type) {
	case int:
		N = num.Repeat[num.DType](num.DType(v), this.Len())
	case []num.DType:
		N = num.Align(v, num.NaN(), this.Len())
	case Series:
		vs := v.DTypes()
		N = num.Align(vs, num.NaN(), this.Len())
	default:
		panic(num.ErrInvalidWindow)
	}
	w := RollingAndExpandingMixin{
		Window: num.Window[num.DType]{V: N},
		Series: this,
	}
	return w
}
