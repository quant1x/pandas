package winpooh32

import "gitee.com/quant1x/pandas/winpooh32/math"

type AlphaType int

const (
	// Specify smoothing factor α directly, 0<α≤1.
	Alpha AlphaType = iota
	// Specify decay in terms of center of mass, α=1/(1+com), for com ≥ 0.
	AlphaCom
	// Specify decay in terms of span, α=2/(span+1), for span ≥ 1.
	AlphaSpan
	// Specify decay in terms of half-life, α=1−exp(−ln(2)/halflife), for halflife > 0.
	AlphaHalflife
)

type ExponentialMovingWindow struct {
	data     Series
	atype    AlphaType
	param    DType
	adjust   bool
	ignoreNA bool
}

func (w ExponentialMovingWindow) Mean() Series {
	var alpha DType

	switch w.atype {
	case Alpha:
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
		alpha = 1 - math.Exp(-math.Ln2/w.param)
	}

	return w.applyMean(w.data.Clone(), alpha)
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
		values []DType = data.Values()
		weight DType   = 1
		last   DType   = 0
	)

	alpha = 1 - alpha
	for t, x := range values {

		w := alpha*weight + 1

		if IsNA(x) {
			if ignoreNA {
				weight = w
			}
			values[t] = last
			continue
		}

		last = last + (x-last)/w

		weight = w

		values[t] = last
	}
}

func (ExponentialMovingWindow) notadjustedMean(data Series, alpha DType, ignoreNA bool) {
	var (
		count  int
		values []DType = data.Values()
		beta   DType   = 1 - alpha
		last   DType   = values[0]
	)
	if IsNA(last) {
		last = 0
		values[0] = last
	}
	for t := 1; t < len(values); t++ {
		x := values[t]

		if IsNA(x) {
			values[t] = last
			continue
		}

		// yt = (1−α)*y(t−1) + α*x(t)
		last = (beta * last) + (alpha * x)
		values[t] = last

		count++
	}
}
