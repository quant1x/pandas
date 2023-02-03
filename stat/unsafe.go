package stat

func math_abs[T StatType](v T) T {
	if v < 0 {
		return v * -1
	}
	return v
}
