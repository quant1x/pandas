package stat

func detectParam[T StatType](v any) (T, []T, error) {
	var base T
	var slice []T
	switch val := v.(type) {
	case []T:
		slice = val
	case T:
		base = val
	}
	return base, slice, nil
}

// AnyToSlice any转切片
func AnyToSlice[T StatType](A any, n int) []T {
	switch v := A.(type) {
	case /*nil, */ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64 /*, bool, string*/ :
		//x := S.EWM(pandas.EW{Alpha: 1 / stat.Any2DType(N), Adjust: false}).Mean().DTypes()
		//return x
		return Repeat[T](v.(T), n)
	case []T:
		return Align(v, T(0), n)
	default:
		panic(ErrUnsupportedType)
	}
}
