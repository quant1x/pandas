package stat

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
	default:
		panic(Throw(x))
	}
	a := self.DTypes()
	s := Div(a, b)
	return NDArray[DType](s)
}
