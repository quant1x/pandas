package formula

import (
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// EMA 指数移动平均,为了精度 S>4*N  EMA至少需要120周期
// alpha=2/(span+1)
// TODO:这个版本是对的, 通达信EMA居然实现了真的序列, 那为啥SMA不是呢?!
func EMA(S stat.Series, N any) any {
	var X []stat.DType
	switch v := N.(type) {
	case int:
		X = stat.Repeat[stat.DType](stat.DType(v), S.Len())
	case stat.Series:
		vs := v.DTypes()
		X = stat.Align(vs, stat.DTypeNaN, S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	k := X[0]
	x := S.EWM(stat.EW{Span: stat.DTypeNaN, Callback: func(idx int) stat.DType {
		j := X[idx]
		if j == 0 {
			j = 1
		} else {
			k = j
		}
		return stat.DType(stat.DType(2) / (j + 1))
	}, Adjust: false}).Mean().Values()
	_ = k
	return x
}

// EMA_v2 通达信公式管理器上提示, EMA(S, N) 相当于SMA(S, N + 1, M=2), 骗子, 根本不对
func EMA_v2(S stat.Series, N any) any {
	M := 2
	var X float32
	switch v := N.(type) {
	case int:
		X = float32(v)
	case stat.Series:
		vs := v.Values()
		fs := stat.SliceToFloat32(vs)
		X = fs[len(fs)-1]
	default:
		panic(exception.New(1, "error window"))
	}
	x := S.EWM(stat.EW{Alpha: float64(M) / float64(X+1), Adjust: false}).Mean().Values()
	return x
}

// EMA_v0 仿SMA实现, 错误
func EMA_v0(S stat.Series, N any) any {
	var X float32
	switch v := N.(type) {
	case int:
		X = float32(v)
	case stat.Series:
		vs := v.Values()
		fs := stat.SliceToFloat32(vs)
		X = fs[len(fs)-1]
	default:
		panic(exception.New(1, "error window"))
	}
	x := S.EWM(stat.EW{Span: stat.DType(X), Adjust: false}).Mean().Values()
	return x
}

// EMA_v1 Rolling(N), 每个都取最后一个, 错误
func EMA_v1(S stat.Series, N any) any {
	x := S.Rolling(N).Apply(func(S stat.Series, N stat.DType) stat.DType {
		r := S.EWM(stat.EW{Span: N, Adjust: false}).Mean().DTypes()
		if len(r) == 0 {
			return stat.DTypeNaN
		}
		return r[len(r)-1]
	}).Values()
	return x

}
