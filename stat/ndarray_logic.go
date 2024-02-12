package stat

import "gitee.com/quant1x/num"

func (self NDArray[T]) Logic(f func(idx int, v any) bool) []bool {
	d := make([]bool, self.Len())
	for i, v := range self {
		d[i] = f(i, v)
	}
	return d
}

func (self NDArray[T]) Eq(x any) Series {
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
	s := num.Equal(a, b)
	return NDArray[bool](s)
}

func (self NDArray[T]) Neq(x any) Series {
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
	s := num.NotEqual(a, b)
	return NDArray[bool](s)
}

func (self NDArray[T]) Gt(x any) Series {
	values := self.Values().([]T)
	bs := num.Gt(values, x)
	return NDArray[bool](bs)
}

func (self NDArray[T]) Gte(x any) Series {
	values := self.Values().([]T)
	bs := num.Gte(values, x)
	return NDArray[bool](bs)
}

func (self NDArray[T]) Lt(x any) Series {
	values := self.Values().([]T)
	bs := num.Lt(values, x)
	return NDArray[bool](bs)
}

func (self NDArray[T]) Lte(x any) Series {
	values := self.Values().([]T)
	bs := num.Lte(values, x)
	return NDArray[bool](bs)
}

func (self NDArray[T]) And(x any) Series {
	values := self.Values().([]T)
	bs := num.And(values, x)
	return NDArray[bool](bs)
}

func (self NDArray[T]) Or(x any) Series {
	values := self.Values().([]T)
	bs := num.Or(values, x)
	return NDArray[bool](bs)
}

func (self NDArray[T]) Not() Series {
	values := self.Values().([]T)
	bs := num.Not(values)
	return NDArray[bool](bs)
}
