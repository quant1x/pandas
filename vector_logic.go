package pandas

import "gitee.com/quant1x/num"

func (this vector[T]) Logic(f func(idx int, v any) bool) []bool {
	d := make([]bool, this.Len())
	for i, v := range this {
		d[i] = f(i, v)
	}
	return d
}

func (this vector[T]) Eq(x any) Series {
	length := this.Len()
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
	a := this.DTypes()
	s := num.Equal(a, b)
	return vector[bool](s)
}

func (this vector[T]) Neq(x any) Series {
	length := this.Len()
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
	a := this.DTypes()
	s := num.NotEqual(a, b)
	return vector[bool](s)
}

func (this vector[T]) Gt(x any) Series {
	values := this.Values().([]T)
	bs := num.Gt(values, x)
	return vector[bool](bs)
}

func (this vector[T]) Gte(x any) Series {
	values := this.Values().([]T)
	bs := num.Gte(values, x)
	return vector[bool](bs)
}

func (this vector[T]) Lt(x any) Series {
	values := this.Values().([]T)
	bs := num.Lt(values, x)
	return vector[bool](bs)
}

func (this vector[T]) Lte(x any) Series {
	values := this.Values().([]T)
	bs := num.Lte(values, x)
	return vector[bool](bs)
}

func (this vector[T]) And(x any) Series {
	values := this.Values().([]T)
	bs := num.And(values, x)
	return vector[bool](bs)
}

func (this vector[T]) Or(x any) Series {
	values := this.Values().([]T)
	bs := num.Or(values, x)
	return vector[bool](bs)
}

func (this vector[T]) Not() Series {
	values := this.Values().([]T)
	bs := num.Not(values)
	return vector[bool](bs)
}
