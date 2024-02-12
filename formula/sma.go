package formula

import (
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// SMA 中国式的SMA,至少需要120周期才精确 (雪球180周期)    alpha=1/(1+com)
func SMA(S stat.Series, N any, M int) stat.Series {
	if M == 0 {
		M = 1
	}
	var X float32
	switch v := N.(type) {
	case int:
		X = float32(v)
	case stat.Series:
		vs := v.Values()
		fs := num.SliceToFloat32(vs)
		X = fs[len(fs)-1]
	default:
		panic(exception.New(1, "error window"))
	}
	//x := S.EWM(stat.EW{Alpha: float64(M) / float64(X), Adjust: false}).Mean().Values()
	x := S.EWM(stat.EW{Alpha: float64(M) / float64(X), Adjust: false}).Mean()
	return x
}

// 最接近
func SMA_v5(S stat.Series, N any, M int) any {
	if M == 0 {
		M = 1
	}
	var X []float32
	switch v := N.(type) {
	case int:
		X = num.Repeat[float32](float32(v), S.Len())
	case stat.Series:
		vs := v.Values()
		X = num.SliceToFloat32(vs)
		X = num.Align(X, num.Nil2Float32, S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	k := X[0]
	x := S.EWM(stat.EW{Alpha: num.Nil2Float64, Callback: func(idx int) num.DType {
		j := X[idx]
		if j == 0 {
			j = 1
		} else {
			k = j
		}
		return float64(M) / float64(j)
	}, Adjust: false}).Mean().Values()
	_ = k
	return x
}

// SMA_v4 听说SMA(S, N, 1) 其实就是MA(S,N), 试验后发现是骗子
func SMA_v4(S stat.Series, N any, M int) any {
	var X []float32
	switch v := N.(type) {
	case int:
		X = num.Repeat[float32](float32(v), S.Len())
	case stat.Series:
		vs := v.Values()
		X = num.SliceToFloat32(vs)
		X = num.Align(X, num.Nil2Float32, S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	return S.Rolling(X).Mean().Values()
}

// SMA_v3 使用滑动窗口
func SMA_v3(S stat.Series, N any, M int) any {
	if M == 0 {
		M = 1
	}
	x := S.Rolling(N).Apply(func(S stat.Series, N num.DType) num.DType {
		r := S.EWM(stat.EW{Alpha: float64(M) / float64(N), Adjust: false}).Mean().Values().([]float64)
		if len(r) == 0 {
			return num.DTypeNaN
		}
		return num.DType(r[len(r)-1])
	}).Values()
	return x
}

// SMA_v1 最原始的python写法
func SMA_v1(S stat.Series, N int, M int) any {
	if M == 0 {
		M = 1
	}
	x := S.EWM(stat.EW{Alpha: float64(M) / float64(N), Adjust: false}).Mean().Values()
	return x
}
