package pandas

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

func (this *NDFrame) Add(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.And(x)
}

func (this *NDFrame) Sub(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Sub(x)
}

func (this *NDFrame) Mul(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Mul(x)
}

func (this *NDFrame) Div(x any) stat.Series {
	vs := this.DTypes()
	s := stat.NDArray[num.DType](vs)
	return s.Div(x)
}
