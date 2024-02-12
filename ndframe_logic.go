package pandas

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

func (this *NDFrame) And(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.And(x)
}

func (this *NDFrame) Eq(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Eq(x)
}

func (this *NDFrame) Neq(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Neq(x)
}

func (this *NDFrame) Gt(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Gt(x)
}

func (this *NDFrame) Gte(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Gte(x)
}

func (this *NDFrame) Lt(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Lt(x)
}

func (this *NDFrame) Lte(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Lte(x)
}

func (this *NDFrame) Or(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Or(x)
}

func (this *NDFrame) Not() stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Not()
}
