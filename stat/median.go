package stat

// Median returns median value of series.
// Linear interpolation is used for odd length.
// TODO:未加验证, 未加速
func Median[T StatType](values []T) DType {
	if len(values) == 0 {
		return DTypeNaN
	}

	if len(values) == 1 {
		return DType(0)
	}

	if len(values)%2 == 0 {
		i := len(values) / 2
		return DType(values[i-1]+values[i]) / 2
	}

	return DType(values[len(values)/2])
}
