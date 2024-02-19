package pandas

import "gitee.com/quant1x/num"

func (this vector[T]) Rolling(param any) RollingAndExpandingMixin {
	window := num.Window[num.DType]{C: num.DTypeNaN}
	switch v := param.(type) {
	case int:
		//N = num.Repeat[num.DType](num.DType(v), this.Len())
		window.C = num.DType(v)
	case []num.DType:
		//N = num.Align(v, num.DTypeNaN, this.Len())
		window.V = v
	case Series:
		vs := v.DTypes()
		//N = num.Align(vs, num.DTypeNaN, this.Len())
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

//func (this vector[T]) Rolling(param any) RollingAndExpandingMixin {
//	var N []num.DType
//	switch v := param.(type) {
//	case int:
//		N = num.Repeat[num.DType](num.DType(v), this.Len())
//	case []num.DType:
//		N = num.Align(v, num.DTypeNaN, this.Len())
//	case Series:
//		vs := v.DTypes()
//		N = num.Align(vs, num.DTypeNaN, this.Len())
//	default:
//		panic(num.ErrInvalidWindow)
//	}
//	w := RollingAndExpandingMixin{
//		Window: N,
//		Series: this,
//	}
//	return w
//}
