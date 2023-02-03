package stat

// Abs 泛型绝对值
func Abs[T StatType](x []T) []T {
	xlen := len(x)
	d := make([]T, xlen)
	for i := 0; i < xlen; i++ {
		if x[i] < 0 {
			d[i] = -1 * (x[i])
		} else {
			d[i] = x[i]
		}
	}
	return d
}
