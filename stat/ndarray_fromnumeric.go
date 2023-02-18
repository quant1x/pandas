package stat

import "github.com/viterin/vek"

func (self NDArray[T]) ArgMax() int {
	return ArgMax2(self)
}

func (self NDArray[T]) ArgMin() int {
	return ArgMin2(self)
}

func (self NDArray[T]) Add(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := Add(a, b)
	return NDArray[DType](s)
}

func (self NDArray[T]) Sub(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := Sub(a, b)
	return NDArray[DType](s)
}

func (self NDArray[T]) Mul(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := Mul(a, b)
	return NDArray[DType](s)
}

func (self NDArray[T]) Div(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := Div(a, b)
	return NDArray[DType](s)
}

func (self NDArray[T]) Eq(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := Equal(a, b)
	return NDArray[bool](s)
}

func (self NDArray[T]) Gt(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := vek.Gt(a, b)
	return NDArray[bool](s)
}

func (self NDArray[T]) Gte(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := vek.Gte(a, b)
	return NDArray[bool](s)
}

func (self NDArray[T]) Lt(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := vek.Lt(a, b)
	return NDArray[bool](s)
}

func (self NDArray[T]) Lte(x any) Series {
	length := self.Len()
	var b []DType
	switch sx := x.(type) {
	case Series:
		b = sx.DTypes()
	case int:
		b = Repeat[DType](DType(sx), length)
	case DType:
		b = Repeat[DType](sx, length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = Align[DType](sx, DTypeNaN, length)
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := vek.Lte(a, b)
	return NDArray[bool](s)
}

func (self NDArray[T]) And(x any) Series {
	length := self.Len()
	var b []bool
	switch sx := x.(type) {
	case Series:
		b = ToBool(sx)
	case int:
		b = Repeat[bool](integer2Bool(sx), length)
	case DType:
		b = Repeat[bool](integer2Bool(sx), length)
	//case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr, float32, float64:
	//	b = Repeat[DType](DType(sx), length)
	case []DType:
		b = __NumberToBool_S(sx)
	default:
		panic(Throw(x))
	}
	a := ToBool(self)
	s := vek.And(a, b)
	return NDArray[bool](s)
}
