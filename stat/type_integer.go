package stat

func integer2Bool[T Number](v T) bool {
	if v != 0 {
		return true
	}
	return false
}
