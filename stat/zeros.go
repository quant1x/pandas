package stat

// Zeros Return a new array of given shape and type, filled with zeros.
//
//	args[0] dtype 基础数据类型
func Zeros[T BaseType](shape int) []T {
	var t T
	return Repeat(t, shape)
}
