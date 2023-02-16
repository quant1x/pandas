package pandas

import "gitee.com/quant1x/pandas/stat"

func (self *NDFrame) Sum() stat.DType {
	fs := self.DTypes()
	return stat.Sum(fs)
}
