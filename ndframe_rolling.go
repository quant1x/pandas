package pandas

import (
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// Rolling RollingAndExpandingMixin
func (this *NDFrame) Rolling(param any) stat.RollingAndExpandingMixin {
	var N []num.DType
	switch v := param.(type) {
	case int:
		N = num.Repeat[num.DType](num.DType(v), this.Len())
	case []num.DType:
		N = num.Align(v, num.DTypeNaN, this.Len())
	case stat.Series:
		vs := v.DTypes()
		N = num.Align(vs, num.DTypeNaN, this.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	w := stat.RollingAndExpandingMixin{
		Window: N,
		Series: this,
	}
	return w
}
