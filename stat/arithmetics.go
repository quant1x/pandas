package stat

// 四则运算 (arithmetics)

// 一元运算 unary operations
func unaryOperations[T Number](x []T, f32 func([]float32) []float32, f64 func([]float64) []float64, cany func([]T) []T) []T {
	var t []T
	if len(x) == 0 {
		return t
	}
	var d any
	var s any
	s = x
	switch fs := s.(type) {
	case []float32:
		d = f32(fs)
	case []float64:
		d = f64(fs)
	default:
		d = cany(x)
	}
	return d.([]T)
}

func unaryOperations1[T Number](x []T, f32 func([]float32) float32, f64 func([]float64) float64, cany func([]T) T) T {
	var t T
	if len(x) == 0 {
		return t
	}
	var d any
	var s any
	s = x
	switch fs := s.(type) {
	case []float32:
		d = f32(fs)
	case []float64:
		d = f64(fs)
	default:
		d = cany(x)
	}
	return d.(T)
}

// 一元运算 unary operations
//
//	运算和返回值是两种类型
func unaryOperations2[T Number, E Number](x []T, f32 func([]float32) E, f64 func([]float64) E, cany func([]T) E) E {
	if len(x) == 0 {
		return E(0)
	}
	var d any
	var s any
	s = x
	switch fs := s.(type) {
	case []float32:
		d = f32(fs)
	case []float64:
		d = f64(fs)
	default:
		d = cany(x)
	}
	return d.(E)
}

// 二元运算 binary operations
//
//	Binary operation
//	calculate
func binaryOperations[T Number](x []T, y any, f32 func(x, y []float32) []float32, f64 func(x, y []float64) []float64, cany func(x, y []T) []T) []T {
	var d any
	xLen := len(x)
	var s any = x
	switch vs := s.(type) {
	case []float32:
		f32s := AnyToSlice[float32](y, xLen)
		d = f32(vs, f32s)
	case []float64:
		f64s := AnyToSlice[float64](y, xLen)
		d = f64(vs, f64s)
	default:
		//panic(ErrUnsupportedType)
		ts := AnyToSlice[T](y, xLen)
		d = cany(x, ts)
	}
	return d.([]T)
}

// 三元运算 triple operations
