package formula

import (
	"github.com/quant1x/num"
	"github.com/quant1x/pandas"
	"github.com/quant1x/x/exception"
)

// SMA 中国式的SMA,至少需要120周期才精确 (雪球180周期)    alpha=1/(1+com)
func SMA(S pandas.Series, N any, M int) pandas.Series {
	return v1SMA(S, N, M)
}

func v1SMA(S pandas.Series, N any, M int) pandas.Series {
	if M == 0 {
		M = 1
	}
	var X float32
	switch v := N.(type) {
	case int:
		X = float32(v)
	case pandas.Series:
		vs := v.Values()
		fs := num.SliceToFloat32(vs)
		X = fs[len(fs)-1]
	default:
		panic(num.ErrInvalidWindow)
	}
	//x := S.EWM(stat.EW{Alpha: float64(M) / float64(X), Adjust: false}).Mean().Values()
	x := S.EWM(pandas.EW{Alpha: float64(M) / float64(X), Adjust: false}).Mean()
	return x
}

func v2SMA(S pandas.Series, N any, M int) pandas.Series {
	panic("not implemented")
}

// v3SMA 使用滑动窗口
func v3SMA(S pandas.Series, N any, M int) any {
	if M == 0 {
		M = 1
	}
	x := S.Rolling(N).Apply(func(S pandas.Series, N num.DType) num.DType {
		r := S.EWM(pandas.EW{Alpha: float64(M) / float64(N), Adjust: false}).Mean().Values().([]float64)
		if len(r) == 0 {
			return num.NaN()
		}
		return num.DType(r[len(r)-1])
	}).Values()
	return x
}

// v4SMA 听说SMA(S, N, 1) 其实就是MA(S,N), 试验后发现是骗子
func v4SMA(S pandas.Series, N any, M int) any {
	var X []float32
	switch v := N.(type) {
	case int:
		X = num.Repeat[float32](float32(v), S.Len())
	case pandas.Series:
		vs := v.Values()
		X = num.SliceToFloat32(vs)
		X = num.Align(X, num.Float32NaN(), S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	return S.Rolling(X).Mean().Values()
}

// 最接近
func v5SMA(S pandas.Series, N any, M int) any {
	if M == 0 {
		M = 1
	}
	var X []float32
	switch v := N.(type) {
	case int:
		X = num.Repeat[float32](float32(v), S.Len())
	case pandas.Series:
		vs := v.Values()
		X = num.SliceToFloat32(vs)
		X = num.Align(X, num.Float32NaN(), S.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	k := X[0]
	x := S.EWM(pandas.EW{Alpha: num.Float64NaN(), Callback: func(idx int) num.DType {
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

// v6SMA 最原始的python写法
func v6SMA(S pandas.Series, N int, M int) any {
	if M == 0 {
		M = 1
	}
	x := S.EWM(pandas.EW{Alpha: float64(M) / float64(N), Adjust: false}).Mean().Values()
	return x
}
