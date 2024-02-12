package pandas

import "gitee.com/quant1x/num"

func (self NDArray[T]) ArgMax() int {
	return num.ArgMax2(self)
}

func (self NDArray[T]) ArgMin() int {
	return num.ArgMin2(self)
}

func (self NDArray[T]) Add(x any) Series {
	length := self.Len()
	var b []num.DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = num.Repeat[num.DType](num.DType(sx), length)
	case num.DType:
		b = num.Repeat[num.DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []num.DType:
		b = num.Align[num.DType](sx, num.DTypeNaN, length)
	default:
		panic(num.TypeError(x))
	}
	a := self.DTypes()
	s := num.Add(a, b)
	return NDArray[num.DType](s)
}

func (self NDArray[T]) Sub(x any) Series {
	length := self.Len()
	var b []num.DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = num.Repeat[num.DType](num.DType(sx), length)
	case num.DType:
		b = num.Repeat[num.DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []num.DType:
		b = num.Align[num.DType](sx, num.DTypeNaN, length)
	default:
		panic(num.TypeError(x))
	}
	a := self.DTypes()
	s := num.Sub(a, b)
	return NDArray[num.DType](s)
}

func (self NDArray[T]) Mul(x any) Series {
	length := self.Len()
	var b []num.DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = num.Repeat[num.DType](num.DType(sx), length)
	case num.DType:
		b = num.Repeat[num.DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []num.DType:
		b = num.Align[num.DType](sx, num.DTypeNaN, length)
	default:
		panic(num.TypeError(x))
	}
	a := self.DTypes()
	s := num.Mul(a, b)
	return NDArray[num.DType](s)
}

func (self NDArray[T]) Div(x any) Series {
	length := self.Len()
	var b []num.DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = num.Repeat[num.DType](num.DType(sx), length)
	case num.DType:
		b = num.Repeat[num.DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []num.DType:
		b = num.Align[num.DType](sx, num.DTypeNaN, length)
	default:
		panic(num.TypeError(x))
	}
	a := self.DTypes()
	s := num.Div(a, b)
	return NDArray[num.DType](s)
}
