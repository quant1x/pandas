package pandas

import "gitee.com/quant1x/pandas/stat"

func (self *NDFrame) Sum() float64 {
	fs := ToFloat64(self)
	return stat.Sum(fs)
}
