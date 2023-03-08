package pandas

import "gitee.com/quant1x/pandas/stat"

func (self *NDFrame) And(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.And(x)
}

func (self *NDFrame) Eq(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Eq(x)
}

func (self *NDFrame) Gt(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Gt(x)
}

func (self *NDFrame) Gte(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Gte(x)
}

func (self *NDFrame) Lt(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Lt(x)
}

func (self *NDFrame) Lte(x any) stat.Series {
	vs := self.DTypes()
	s := stat.NDArray[stat.DType](vs)
	return s.Lte(x)
}
