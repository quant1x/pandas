package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"math"
)

type AlphaType int

// https://pandas.pydata.org/pandas-docs/stable/reference/api/pandas.DataFrame.ewm.html
const (
	// AlphaAlpha Specify smoothing factor α directly, 0<α≤1.
	AlphaAlpha AlphaType = iota
	// AlphaCom Specify decay in terms of center of mass, α=1/(1+com), for com ≥ 0.
	AlphaCom
	// AlphaSpan Specify decay in terms of span, α=2/(span+1), for span ≥ 1.
	AlphaSpan
	// AlphaHalfLife Specify decay in terms of half-life, α=1−exp(−ln(2)/halflife), for halflife > 0.
	AlphaHalfLife
)

// EW (Factor) 指数加权(EW)计算Alpha 结构属性非0即为有效启动同名算法
type EW struct {
	Com      stat.DType // 根据质心指定衰减
	Span     stat.DType // 根据跨度指定衰减
	HalfLife stat.DType // 根据半衰期指定衰减
	Alpha    stat.DType // 直接指定的平滑因子α
	Adjust   bool       // 除以期初的衰减调整系数以核算 相对权重的不平衡（将 EWMA 视为移动平均线）
	IgnoreNA bool       // 计算权重时忽略缺失值
	Callback func(idx int) stat.DType
}

// ExponentialMovingWindow 加权移动窗口
type ExponentialMovingWindow struct {
	data       Series     // 序列
	atype      AlphaType  // 计算方式: com/span/halflefe/alpha
	param      stat.DType // 参数类型为浮点
	adjust     bool       // 默认为真, 是否调整, 默认真时, 计算序列的EW移动平均线, 为假时, 计算指数加权递归
	ignoreNA   bool       // 默认为假, 计算权重时是否忽略缺失值NaN
	minPeriods int        // 默认为0, 窗口中具有值所需的最小观测值数,否则结果为NaN
	axis       int        // {0,1}, 默认为0, 0跨行计算, 1跨列计算
	cb         func(idx int) stat.DType
}

// EWM provides exponential weighted calculations.
func (s *NDFrame) EWM(alpha EW) ExponentialMovingWindow {
	atype := AlphaAlpha
	param := 0.00
	adjust := alpha.Adjust
	ignoreNA := alpha.IgnoreNA
	if alpha.Com != 0 {
		atype = AlphaCom
		param = alpha.Com
	} else if alpha.Span != 0 {
		atype = AlphaSpan
		param = alpha.Span
	} else if alpha.HalfLife != 0 {
		atype = AlphaHalfLife
		param = alpha.HalfLife
	} else {
		atype = AlphaAlpha
		param = alpha.Alpha
	}

	dest := NewSeries(SERIES_TYPE_FLOAT64, s.name, s.Values())
	return ExponentialMovingWindow{
		data:     dest,
		atype:    atype,
		param:    param,
		adjust:   adjust,
		ignoreNA: ignoreNA,
		cb:       alpha.Callback,
	}
}

func (w ExponentialMovingWindow) Mean() Series {
	var alpha stat.DType

	switch w.atype {
	case AlphaAlpha:
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

	case AlphaHalfLife:
		if w.param <= 0 {
			panic("halflife param must be > 0")
		}
		alpha = 1 - math.Exp(-math.Ln2/w.param)
	}

	return w.applyMean(w.data, alpha)
}

func (w ExponentialMovingWindow) applyMean(data Series, alpha stat.DType) Series {
	if w.adjust {
		w.adjustedMean(data, alpha, w.ignoreNA)
	} else {
		w.notadjustedMean(data, alpha, w.ignoreNA)
	}
	return data
}

func (w ExponentialMovingWindow) adjustedMean(data Series, alpha stat.DType, ignoreNA bool) {
	var (
		values            = data.Values().([]stat.DType)
		weight stat.DType = 1
		last              = values[0]
	)

	alpha = 1 - alpha
	for t := 1; t < len(values); t++ {

		w := alpha*weight + 1
		x := values[t]
		if stat.DTypeIsNaN(x) {
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

func (w ExponentialMovingWindow) notadjustedMean(data Series, alpha stat.DType, ignoreNA bool) {
	hasCallback := false
	if stat.DTypeIsNaN(alpha) {
		hasCallback = true
		alpha = w.cb(0)
	}
	var (
		count  int
		values = data.Values().([]stat.DType)
		//values = data.DTypes() // Dtypes有复制功能
		beta = 1 - alpha
		last = values[0]
	)
	if Float64IsNaN(last) {
		last = 0
		values[0] = last
	}
	for t := 1; t < len(values); t++ {
		x := values[t]

		if stat.DTypeIsNaN(x) {
			values[t] = last
			continue
		}
		if hasCallback {
			alpha = w.cb(t)
			beta = 1 - alpha
		}
		// yt = (1−α)*y(t−1) + α*x(t)
		last = (beta * last) + (alpha * x)
		if stat.DTypeIsNaN(last) {
			last = values[t-1]
		}
		values[t] = last

		count++
	}
}
