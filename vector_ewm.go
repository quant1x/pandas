package pandas

import "gitee.com/quant1x/num"

func (this Vector[T]) EWM(alpha EW) ExponentialMovingWindow {
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

	dest := ToSeries[num.DType]()
	dest = dest.Append(this)
	return ExponentialMovingWindow{
		Data:     dest,
		AType:    atype,
		Param:    param,
		Adjust:   adjust,
		IgnoreNA: ignoreNA,
		Cb:       alpha.Callback,
	}
}
