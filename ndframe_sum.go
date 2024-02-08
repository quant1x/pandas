package pandas

import "gitee.com/quant1x/pandas/stat"

func (this *NDFrame) Sum() stat.DType {
	fs := this.DTypes()
	return stat.Sum(fs)
}
