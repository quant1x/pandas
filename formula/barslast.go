package formula

import (
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
)

// 为了测试SMA,BARSLAST必须要先实现, 给SMA提供序列换参数, 以便验证, python那边还没实现
func barslast(S pandas.Series) []float32 {
	fs := pandas.ToFloat32(S)
	as := stat.Repeat[float32](1, S.Len())
	bs := stat.Repeat[float32](0, S.Len())
	x := stat.Where(fs, as, bs)
	M := []float32{0}
	M = append(M, x...)
	for i := 1; i < len(M); i++ {
		if int(M[i]) != 0 {
			M[i] = 0
		} else {
			M[i] = M[i-1] + 1
		}
		//fmt.Println(M[i])
	}
	return M[1:]
}
