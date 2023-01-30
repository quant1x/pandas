package pandas

import (
	"gitee.com/quant1x/pandas/algorithms/winpooh32"
	math2 "gitee.com/quant1x/pandas/algorithms/winpooh32/math"
)

type DType = float64

type AlphaType int

// https://pandas.pydata.org/pandas-docs/stable/reference/api/pandas.DataFrame.ewm.html
const (
	// Specify smoothing factor α directly, 0<α≤1.
	AlphaNil AlphaType = iota
	// Specify decay in terms of center of mass, α=1/(1+com), for com ≥ 0.
	AlphaCom
	// Specify decay in terms of span, α=2/(span+1), for span ≥ 1.
	AlphaSpan
	// Specify decay in terms of half-life, α=1−exp(−ln(2)/halflife), for halflife > 0.
	AlphaHalflife
)

// EW(Factor) 指数加权(EW)计算Alpha 结构属性非0即为有效启动同名算法
type EW struct {
	Com      float64 // 根据质心指定衰减
	Span     float64 // 根据跨度指定衰减
	Halflife float64 // 根据半衰期指定衰减
	Alpha    float64 // 直接指定的平滑因子α
	Adjust   bool    // 除以期初的衰减调整系数以核算 相对权重的不平衡（将 EWMA 视为移动平均线）
	IgnoreNA bool    // 计算权重时忽略缺失值
}

type ExponentialMovingWindow struct {
	data       Series    // 序列
	atype      AlphaType // 计算方式: com/span/halflefe/alpha
	param      DType     // 参数类型为浮点
	adjust     bool      // 默认为真, 是否调整, 默认真时, 计算序列的EW移动平均线, 为假时, 计算指数加权递归
	ignoreNA   bool      // 默认为假, 计算权重时是否忽略缺失值NaN
	minPeriods int       // 默认为0, 窗口中具有值所需的最小观测值数,否则结果为NaN
	axis       int       // {0,1}, 默认为0, 0跨行计算, 1跨列计算
}

// EWM provides exponential weighted calculations.
func (s *SeriesFloat64) EWM(alpha EW) ExponentialMovingWindow {
	atype := AlphaNil
	param := 0.00
	adjust := alpha.Adjust
	ignoreNA := alpha.IgnoreNA
	if alpha.Com != 0 {
		atype = AlphaCom
		param = alpha.Com
	} else if alpha.Span != 0 {
		atype = AlphaSpan
		param = alpha.Span
	} else if alpha.Halflife != 0 {
		atype = AlphaHalflife
		param = alpha.Halflife
	} else {
		atype = AlphaNil
		param = alpha.Alpha
	}

	dest := NewSeriesFloat64(s.name, s.Values())
	return ExponentialMovingWindow{
		data:     dest,
		atype:    atype,
		param:    param,
		adjust:   adjust,
		ignoreNA: ignoreNA,
	}
}

func (w ExponentialMovingWindow) Mean() Series {
	var alpha DType

	switch w.atype {
	case AlphaNil:
		if w.param <= 0 {
			panic("alpha param must be > 0")
		}
		alpha = w.param

	case AlphaCom:
		if w.param <= 0 {
			panic("com param must be >= 0")
		}
		alpha = 1 / (1 + w.param)

	case AlphaSpan:
		if w.param < 1 {
			panic("span param must be >= 1")
		}
		alpha = 2 / (w.param + 1)

	case AlphaHalflife:
		if w.param <= 0 {
			panic("halflife param must be > 0")
		}
		alpha = 1 - math2.Exp(-math2.Ln2/w.param)
	}

	return w.applyMean(w.data, alpha)
}

func (w ExponentialMovingWindow) applyMean(data Series, alpha DType) Series {
	if w.adjust {
		w.adjustedMean(data, alpha, w.ignoreNA)
	} else {
		w.notadjustedMean(data, alpha, w.ignoreNA)
	}
	return data
}

func (ExponentialMovingWindow) adjustedMean(data Series, alpha DType, ignoreNA bool) {
	var (
		values       = data.Values().([]float64)
		weight DType = 1
		last         = values[0]
	)

	alpha = 1 - alpha
	for t := 1; t < len(values); t++ {

		w := alpha*weight + 1
		x := values[t]
		if winpooh32.IsNA(x) {
			if ignoreNA {
				weight = w
			}
			values[t] = last
			continue
		}

		last = last + (x-last)/(w)
		weight = w
		values[t] = last
	}
}

func (ExponentialMovingWindow) notadjustedMean(data Series, alpha DType, ignoreNA bool) {
	var (
		count  int
		values = data.Values().([]float64)
		beta   = 1 - alpha
		last   = values[0]
	)
	if winpooh32.IsNA(last) {
		last = 0
		values[0] = last
	}
	for t := 1; t < len(values); t++ {
		x := values[t]

		if winpooh32.IsNA(x) {
			values[t] = last
			continue
		}

		// yt = (1−α)*y(t−1) + α*x(t)
		last = (beta * last) + (alpha * x)
		values[t] = last

		count++
	}
}
