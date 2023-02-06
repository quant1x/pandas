package stat

import "github.com/viterin/vek"

// 初始化 avx2
// 可以参考另一个实现库 gonum.org/v1/gonum/stat
func init() {
	// 开启加速选项
	vek.SetAcceleration(true)
}
