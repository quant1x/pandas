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
	length := len(x)
	var s any = x
	switch vs := s.(type) {
	case []float32:
		f32s := AnyToSlice[float32](y, length)
		d = f32(vs, f32s)
	case []float64:
		f64s := AnyToSlice[float64](y, length)
		d = f64(vs, f64s)
	default:
		ys := AnyToSlice[T](y, length)
		d = cany(x, ys)
	}
	return d.([]T)
}

func binaryOperations2[T BaseType, E BaseType](x, y []T, f32 func(x, y []float32) []E, f64 func(x, y []float64) []E, cany func(x, y []T) []E) []E {
	var d any
	length := len(x)
	var s any = x
	switch vs := s.(type) {
	case []float32:
		f32s := AnyToSlice[float32](y, length)
		d = f32(vs, f32s)
	case []float64:
		f64s := AnyToSlice[float64](y, length)
		d = f64(vs, f64s)
	default:
		ys := AnyToSlice[T](y, length)
		d = cany(x, ys)
	}
	return d.([]E)
}

// 三元运算 triple operations
/*
const (
	__k_calc_add = 1 // 加
	__k_calc_sub = 2 // 减
	__k_calc_mul = 3 // 乘
	__k_calc_div = 4 // 除
	__k_calc_mod = 5 // 取模
)

var (
	// 加
	__calc_add = func(f1, f2 DType) DType {
		return f1 + f2
	}
	// 减
	__calc_sub = func(f1, f2 DType) DType {
		return f1 - f2
	}
	// 乘
	__calc_mul = func(f1, f2 DType) DType {
		return f1 * f2
	}
	// 除
	__calc_div = func(f1, f2 DType) DType {
		return f1 / f2
	}
	// 取模
	__calc_mod = func(f1, f2 DType) DType {
		return math.Mod(f1, f2)
	}
)

// 重构

func __arithmetic[T ~[]E, E Number](x T, y any, c int, calculator func(f1, f2 DType) E) []E {
	if __y, ok := y.(Series); ok {
		y = __y.Values()
	}
	var d = []E{}
	switch Y := y.(type) {
	case nil, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
		f2 := Any2DType(Y)
		d = __arithmetic_dtype(x, f2, c, calculator)
	case []float32:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []float64:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []int:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []int8:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []int16:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []int32:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []int64:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []uint:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []uint8:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []uint16:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []uint32:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []uint64:
		d = __arithmetic_slice(x, Y, c, calculator)
	case []uintptr:
		d = __arithmetic_slice(x, Y, c, calculator)
	//case []string:
	//	d = __arithmetic_slice(x, Y, c, calculator)
	//case []bool:
	//	d = __arithmetic_slice(x, Y, c, calculator)
	default:
		// 其它未知类型抛异常
		panic(Throw(y))
	}
	return d
}

// 切片和dtype对比, 不用考虑slice长度对齐的问题
func __arithmetic_dtype[T ~[]E, E Number](x T, y DType, c int, calculator func(f1, f2 DType) E) []E {
	var d any
	xLen := len(x)
	kind := checkoutRawType(x)
	switch {
	case kind == SERIES_TYPE_FLOAT64 && c == __k_calc_add:
		fs := make([]float64, xLen)
		d = num.AddNumber_Into(fs, any(x).([]float64), y)
	case kind == SERIES_TYPE_FLOAT64 && c == __k_calc_sub:
		fs := make([]float64, xLen)
		d = num.SubNumber_Into(fs, any(x).([]float64), y)
	case kind == SERIES_TYPE_FLOAT64 && c == __k_calc_mul:
		fs := make([]float64, xLen)
		d = num.MulNumber_Into(fs, any(x).([]float64), y)
	case kind == SERIES_TYPE_FLOAT64 && c == __k_calc_div:
		fs := make([]float64, xLen)
		d = num.DivNumber_Into(fs, any(x).([]float64), y)
	case kind == SERIES_TYPE_FLOAT32 && c == __k_calc_add:
		fs := make([]float32, xLen)
		d = num32.AddNumber_Into(fs, any(x).([]float32), float32(y))
	case kind == SERIES_TYPE_FLOAT32 && c == __k_calc_sub:
		fs := make([]float32, xLen)
		d = num32.SubNumber_Into(fs, any(x).([]float32), float32(y))
	case kind == SERIES_TYPE_FLOAT32 && c == __k_calc_mul:
		fs := make([]float32, xLen)
		d = num32.MulNumber_Into(fs, any(x).([]float32), float32(y))
	case kind == SERIES_TYPE_FLOAT32 && c == __k_calc_div:
		fs := make([]float32, xLen)
		d = num32.DivNumber_Into(fs, any(x).([]float32), float32(y))
	default:
		b := y
		bs := make([]E, xLen)
		for i := 0; i < xLen; i++ {
			a := Any2DType(x[i])
			bs[i] = calculator(a, b)
		}
	}
	return d.([]E)
}

// 切片和切片对比
func __arithmetic_slice[T ~[]E, E Number, T2 ~[]E2, E2 Number](x T, y T2, c int, calculator func(f1, f2 DType) E) []E {
	var d any
	xLen := len(x)
	yLen := len(y)

	if xLen >= yLen {
		es := make([]E, xLen)
		switch xs := any(x).(type) {
		case []float64:
			num.Add_Into(es[:yLen], xs[:yLen], any(y).([]float64)[:yLen])
		}
		switch {
		case xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_calc_add:
			num.Add_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		case xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_calc_sub:
			num.Sub_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		case xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_calc_mul:
			num.Mul_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		case xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_calc_div:
			num.Div_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		case xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_calc_add:
			num.Add_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		case xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_calc_sub:
			num.Sub_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		case xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_calc_mul:
			num.Mul_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		case xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_calc_div:
			num.Div_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		}
		if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_calc_add {
			num.Gt_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		} else if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_compare_gte {
			es := make([]float64, xLen)
			num.Gte_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		} else if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_compare_lt {
			es := make([]float64, xLen)
			num.Lt_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		} else if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_compare_lte {
			es := make([]float64, xLen)
			num.Lte_Into(es[:yLen], any(x).([]float64)[:yLen], any(y).([]float64)[:yLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_gt {
			num32.Gt_Into(es[:yLen], any(x).([]float32)[:yLen], any(y).([]float32)[:yLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_gte {
			num32.Gte_Into(es[:yLen], any(x).([]float32)[:yLen], any(y).([]float32)[:yLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_lt {
			num32.Lt_Into(es[:yLen], any(x).([]float32)[:yLen], any(y).([]float32)[:yLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_lte {
			num32.Lte_Into(es[:yLen], any(x).([]float32)[:yLen], any(y).([]float32)[:yLen])
		} else if xKind == SERIES_TYPE_BOOL && xKind == yKind && c == __k_compare_and {
			num.And_Into(es[:yLen], any(x).([]bool)[:yLen], any(y).([]bool)[:yLen])
		} else {
			for i := 0; i < yLen; i++ {
				f1 := Any2DType(x[i])
				f2 := Any2DType(y[i])
				es[i] = calculator(f1, f2)
			}
		}
		for i := yLen; i < xLen; i++ {
			f1 := Any2DType(x[i])
			f2 := DType(0)
			es[i] = calculator(f1, f2)
		}
	} else {
		es = make([]bool, yLen)
		if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_compare_gt {
			num.Gt_Into(es[:xLen], any(x).([]float64)[:xLen], any(y).([]float64)[:xLen])
		} else if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_compare_gte {
			num.Gte_Into(es[:xLen], any(x).([]float64)[:xLen], any(y).([]float64)[:xLen])
		} else if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_compare_lt {
			num.Lt_Into(es[:xLen], any(x).([]float64)[:xLen], any(y).([]float64)[:xLen])
		} else if xKind == SERIES_TYPE_FLOAT64 && xKind == yKind && c == __k_compare_lte {
			num.Lte_Into(es[:xLen], any(x).([]float64)[:xLen], any(y).([]float64)[:xLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_gt {
			num32.Gt_Into(es[:xLen], any(x).([]float32)[:xLen], any(y).([]float32)[:xLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_gte {
			num32.Gte_Into(es[:xLen], any(x).([]float32)[:xLen], any(y).([]float32)[:xLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_lt {
			num32.Lt_Into(es[:xLen], any(x).([]float32)[:xLen], any(y).([]float32)[:xLen])
		} else if xKind == SERIES_TYPE_FLOAT32 && xKind == yKind && c == __k_compare_lte {
			num32.Lte_Into(es[:xLen], any(x).([]float32)[:xLen], any(y).([]float32)[:xLen])
		} else if xKind == SERIES_TYPE_BOOL && xKind == yKind && c == __k_compare_and {
			num.And_Into(es[:xLen], any(x).([]bool)[:xLen], any(y).([]bool)[:xLen])
		} else {
			for i := 0; i < xLen; i++ {
				f1 := Any2DType(x[i])
				f2 := Any2DType(y[i])
				es[i] = calculator(f1, f2)
			}
		}
		for i := xLen; i < yLen; i++ {
			f1 := DType(0)
			f2 := Any2DType(y[i])
			es[i] = calculator(f1, f2)
		}
	}
	return d.([]E)
}
*/
