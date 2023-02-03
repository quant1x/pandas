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
