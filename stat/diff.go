package stat

// Diff returns the n-th differences of the given array.
// TODO:这个代码有问题, 需要从generic_diff迁移过来
func Diff[T Number](x []T) []T {
	var result []T
	for i := 1; i < len(x); i++ {
		result = append(result, x[i]-x[i-1])
	}
	return result
}
