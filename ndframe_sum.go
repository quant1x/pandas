package pandas

import (
	"gitee.com/quant1x/num"
)

func (this *NDFrame) Sum() num.DType {
	fs := this.DTypes()
	return num.Sum(fs)
}
