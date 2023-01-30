package df

import (
	"gitee.com/quant1x/pandas/algorithms/winpooh32"
	math2 "gitee.com/quant1x/pandas/algorithms/winpooh32/math"
)

type DType = float64

type AlphaType int

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

type Alpha struct {
	Com      float64
	Span     float64
	Halflife float64
	At       AlphaType
}

type ExponentialMovingWindow struct {
	data     Series
	atype    AlphaType
	param    DType
	adjust   bool
	ignoreNA bool
}

// EWM provides exponential weighted calculations.
func (d Series) EWM(alpha Alpha, adjust bool, ignoreNA bool) ExponentialMovingWindow {
	//atype AlphaType, param DType
	atype := alpha.At
	param := 0.00

	switch atype {
	case AlphaNil:
		param = alpha.Com
	case AlphaSpan:
		param = alpha.Span
	case AlphaHalflife:
		param = alpha.Halflife
	}

	dest := NewSeries(d.Float(), Float, d.Name)
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

	//return w.applyMean(w.data.Clone(), alpha)
	return w.applyMean(w.data.Copy(), alpha)
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
		values       = data.elements
		weight DType = 1
		//last   DType = 0
		last DType = values.Elem(0).Float()
	)

	alpha = 1 - alpha
	for t := 1; t < values.Len(); t++ {

		w := alpha*weight + 1
		x := values.Elem(t).Float()
		if winpooh32.IsNA(x) {
			if ignoreNA {
				weight = w
			}
			//values[t] = last
			values.Elem(t).Set(last)
			continue
		}

		last = last + (x-last)/(w)
		weight = w
		values.Elem(t).Set(last)
	}
}

func (ExponentialMovingWindow) notadjustedMean(data Series, alpha DType, ignoreNA bool) {
	var (
		count int
		//values []DType = data.Values()
		//values []DType = data.Float()
		values floatElements = data.elements.(floatElements)

		beta DType = 1 - alpha
		//last DType = values[0]
		last DType = values.Elem(0).Float()
	)
	if winpooh32.IsNA(last) {
		last = 0
		//values[0] = last
		values.Elem(0).Set(last)
	}
	for t := 1; t < values.Len(); t++ {
		//x := values[t]
		x := values.Elem(t).Float()

		if winpooh32.IsNA(x) {
			//values[t] = last
			values.Elem(t).Set(last)
			continue
		}

		// yt = (1−α)*y(t−1) + α*x(t)
		last = (beta * last) + (alpha * x)
		values.Elem(t).Set(last)

		count++
	}
}
