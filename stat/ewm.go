package stat

import "math"

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
	Com      DType // 根据质心指定衰减
	Span     DType // 根据跨度指定衰减
	HalfLife DType // 根据半衰期指定衰减
	Alpha    DType // 直接指定的平滑因子α
	Adjust   bool  // 除以期初的衰减调整系数以核算 相对权重的不平衡（将 EWMA 视为移动平均线）
	IgnoreNA bool  // 计算权重时忽略缺失值
	Callback func(idx int) DType
}

// ExponentialMovingWindow 加权移动窗口
type ExponentialMovingWindow struct {
	Data       Series    // 序列
	AType      AlphaType // 计算方式: com/span/halflefe/alpha
	Param      DType     // 参数类型为浮点
	Adjust     bool      // 默认为真, 是否调整, 默认真时, 计算序列的EW移动平均线, 为假时, 计算指数加权递归
	IgnoreNA   bool      // 默认为假, 计算权重时是否忽略缺失值NaN
	minPeriods int       // 默认为0, 窗口中具有值所需的最小观测值数,否则结果为NaN
	axis       int       // {0,1}, 默认为0, 0跨行计算, 1跨列计算
	Cb         func(idx int) DType
}

func (w ExponentialMovingWindow) Mean() Series {
	var alpha DType

	switch w.AType {
	case AlphaAlpha:
		if w.Param <= 0 {
			panic("alpha param must be > 0")
		}
		alpha = w.Param

	case AlphaCom:
		if w.Param <= 0 {
			panic("com param must be >= 0")
		}
		alpha = 1 / (1 + w.Param)

	case AlphaSpan:
		if w.Param < 1 {
			panic("span param must be >= 1")
		}
		alpha = 2 / (w.Param + 1)

	case AlphaHalfLife:
		if w.Param <= 0 {
			panic("halflife param must be > 0")
		}
		alpha = 1 - math.Exp(-math.Ln2/w.Param)
	}

	return w.applyMean(w.Data, alpha)
}

func (w ExponentialMovingWindow) applyMean(data Series, alpha DType) Series {
	if w.Adjust {
		w.adjustedMean(data, alpha, w.IgnoreNA)
	} else {
		w.notAdjustedMean(data, alpha, w.IgnoreNA)
	}
	return data
}

func (w ExponentialMovingWindow) adjustedMean(data Series, alpha DType, ignoreNA bool) {
	var (
		values       = data.Values().([]DType)
		weight DType = 1
		last         = values[0]
	)

	alpha = 1 - alpha
	for t := 1; t < len(values); t++ {

		w := alpha*weight + 1
		x := values[t]
		if DTypeIsNaN(x) {
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

func (w ExponentialMovingWindow) notAdjustedMean(data Series, alpha DType, ignoreNA bool) {
	hasCallback := false
	if DTypeIsNaN(alpha) {
		hasCallback = true
		alpha = w.Cb(0)
	}
	var (
		count  int
		values = data.Values().([]DType)
		beta   = 1 - alpha
		last   = values[0]
	)
	if DTypeIsNaN(last) {
		last = 0
		values[0] = last
	}
	for t := 1; t < len(values); t++ {
		x := values[t]

		if DTypeIsNaN(x) {
			values[t] = last
			continue
		}
		if hasCallback {
			alpha = w.Cb(t)
			beta = 1 - alpha
		}
		// yt = (1−α)*y(t−1) + α*x(t)
		last = (beta * last) + (alpha * x)
		if DTypeIsNaN(last) {
			last = values[t-1]
		}
		values[t] = last

		count++
	}
}
