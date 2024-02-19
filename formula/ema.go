package formula

import (
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// EMA 指数移动平均,为了精度 S>4*N  EMA至少需要120周期
// alpha=2/(span+1)
// TODO:这个版本是对的, 通达信EMA居然实现了真的序列, 那为啥SMA不是呢?!
func EMA(S pandas.Series, N any) pandas.Series {
	var X []num.DType
	switch v := N.(type) {
	case int:
		X = num.Repeat[num.DType](num.DType(v), S.Len())
	case pandas.Series:
		vs := v.DTypes()
		X = num.Align(vs, num.DTypeNaN, S.Len())
	default:
		panic(num.ErrInvalidWindow)
	}
	k := X[0]
	x := S.EWM(pandas.EW{Span: num.DTypeNaN, Callback: func(idx int) num.DType {
		j := X[idx]
		if j == 0 {
			j = 1
		} else {
			k = j
		}
		return num.DType(num.DType(2) / (j + 1))
	}, Adjust: false}).Mean()
	_ = k
	return x
}

// EMA_v2 通达信公式管理器上提示, EMA(S, N) 相当于SMA(S, N + 1, M=2), 骗子, 根本不对
func EMA_v2(S pandas.Series, N any) any {
	M := 2
	var X float32
	switch v := N.(type) {
	case int:
		X = float32(v)
	case pandas.Series:
		vs := v.Values()
		fs := num.SliceToFloat32(vs)
		X = fs[len(fs)-1]
	default:
		panic(exception.New(1, "error window"))
	}
	x := S.EWM(pandas.EW{Alpha: float64(M) / float64(X+1), Adjust: false}).Mean().Values()
	return x
}

// EMA_v0 仿SMA实现, 错误
func EMA_v0(S pandas.Series, N any) any {
	var X float32
	switch v := N.(type) {
	case int:
		X = float32(v)
	case pandas.Series:
		vs := v.Values()
		fs := num.SliceToFloat32(vs)
		X = fs[len(fs)-1]
	default:
		panic(exception.New(1, "error window"))
	}
	x := S.EWM(pandas.EW{Span: num.DType(X), Adjust: false}).Mean().Values()
	return x
}

// EMA_v1 Rolling(N), 每个都取最后一个, 错误
func EMA_v1(S pandas.Series, N any) any {
	x := S.Rolling(N).Apply(func(S pandas.Series, N num.DType) num.DType {
		r := S.EWM(pandas.EW{Span: N, Adjust: false}).Mean().DTypes()
		if len(r) == 0 {
			return num.DTypeNaN
		}
		return r[len(r)-1]
	}).Values()
	return x

}

// AlphaOfEMA 根据周期是计算α值
//
//	EMA的计算是全部数据, 所以不用考虑第一个元素的情况
func AlphaOfEMA(n int) float64 {
	alpha := float64(2) / float64(1+n)
	return alpha
}

// EmaIncr 增量计算EMA
//
//	通过上一条数值last, alpha和最新值计算
//	yt = (1−α)*y(t−1) + α*x(t)
func EmaIncr(now, last, alpha float64) float64 {
	current := (1-alpha)*last + alpha*now
	return current
}
