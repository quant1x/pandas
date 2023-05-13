package pandas

import (
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/pandas/stat"
)

// Rolling RollingAndExpandingMixin
func (self *NDFrame) Rolling(param any) stat.RollingAndExpandingMixin {
	var N []stat.DType
	switch v := param.(type) {
	case int:
		N = stat.Repeat[stat.DType](stat.DType(v), self.Len())
	case []stat.DType:
		N = stat.Align(v, stat.DTypeNaN, self.Len())
	case stat.Series:
		vs := v.DTypes()
		N = stat.Align(vs, stat.DTypeNaN, self.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	w := stat.RollingAndExpandingMixin{
		Window: N,
		Series: self,
	}
	return w
}
