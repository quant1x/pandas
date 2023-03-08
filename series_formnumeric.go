package pandas

import "gitee.com/quant1x/pandas/stat"

func (self *NDFrame) Add(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.And(x)
}

func (self *NDFrame) Sub(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Sub(x)
}

func (self *NDFrame) Mul(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Mul(x)
}

func (self *NDFrame) Div(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Div(x)
}
