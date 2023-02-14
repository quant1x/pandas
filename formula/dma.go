package formula

import (
	"gitee.com/quant1x/pandas/stat"
)

// DMA 返回动态移动平均
//
//	求S的动态移动平均, A作平滑因子,必须 0<A<1 (此为核心函数，非指标）
//	用法:
//	DMA(X,A),求X的动态移动平均
//	算法:Y=A*X+(1-A)*Y',其中Y'表示上一周期Y值,A必须大于0且小于1.A支持变量
//	例如:
//	DMA(CLOSE,VOL/CAPITAL)表示求以换手率作平滑因子的平均价
func DMA(S stat.Series, A any) []stat.DType {
	switch N := A.(type) {
	case /*nil, */ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64 /*, bool, string*/ :
		x := S.EWM(stat.EW{Alpha: 1 / stat.Any2DType(N), Adjust: false}).Mean().DTypes()
		return x
	case []stat.DType:
		s := S.DTypes()
		stat.Fill(N, 1.0, true)
		Y := stat.Repeat(stat.DType(0), len(s))
		Y[0] = s[0]
		for i := 1; i < len(s); i++ {
			a := 1 / N[i]
			if stat.DTypeIsNaN(a) {
				a = 1
			}
			Y[i] = a*s[i] + (1-a)*Y[i-1]
		}
		return Y
	case stat.Series:
		s := S.DTypes()
		M := N.DTypes()
		stat.Fill(M, 1.0, true)
		Y := stat.Repeat(stat.DType(0), len(s))
		Y[0] = s[0]
		for i := 1; i < len(s); i++ {
			//fmt.Println("M[i] =", M[i])
			a := 1 / M[i]
			if stat.DTypeIsNaN(a) {
				a = 1
			}
			//fmt.Println("a =", a)
			Y[i] = a*s[i] + (1-a)*Y[i-1]
			//fmt.Println(Y[i])
		}
		return Y
	}
	return []stat.DType{}
}
