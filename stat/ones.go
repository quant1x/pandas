package stat

// Ones v -> shape
func Ones[T Number](v []T) []T {
	return Repeat[T](T(1), len(v))
}
