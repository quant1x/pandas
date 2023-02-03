package stat

// Data alignment
func align[T StatType](x []T, a T, dLen int) []T {
	d := []T{}
	xLen := len(x)
	if xLen >= dLen {
		// 截断
		copy(d, x[0:dLen])
	} else {
		// 扩展内存
		d = make([]T, dLen)
		copy(d, x)
		//avx2.RepeatAll(d[xLen:], a)
		for i := xLen; i < dLen; i++ {
			d[i] = a
		}
	}
	return d
}
