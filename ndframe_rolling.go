package pandas

import (
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
)

// Rolling RollingAndExpandingMixin
func (this *NDFrame) Rolling(param any) RollingAndExpandingMixin {
	var N []num.DType
	switch v := param.(type) {
	case int:
		N = num.Repeat[num.DType](num.DType(v), this.Len())
	case []num.DType:
		N = num.Align(v, num.DTypeNaN, this.Len())
	case Series:
		vs := v.DTypes()
		N = num.Align(vs, num.DTypeNaN, this.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	w := RollingAndExpandingMixin{
		Window: N,
		Series: this,
	}
	return w
}
