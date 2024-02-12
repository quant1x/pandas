package pandas

import (
	"gitee.com/quant1x/num"
)

func (this *NDFrame) Add(x any) Series {
	vs := this.DTypes()
	s := NDArray[num.DType](vs)
	return s.And(x)
}

func (this *NDFrame) Sub(x any) Series {
	vs := this.DTypes()
	s := NDArray[num.DType](vs)
	return s.Sub(x)
}

func (this *NDFrame) Mul(x any) Series {
	vs := this.DTypes()
	s := NDArray[num.DType](vs)
	return s.Mul(x)
}

func (this *NDFrame) Div(x any) Series {
	vs := this.DTypes()
	s := NDArray[num.DType](vs)
	return s.Div(x)
}
