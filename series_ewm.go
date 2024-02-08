package pandas

import (
	"gitee.com/quant1x/pandas/stat"
)

// EWM provides exponential weighted calculations.
func (this *NDFrame) EWM(alpha stat.EW) stat.ExponentialMovingWindow {
	atype := stat.AlphaAlpha
	param := 0.00
	adjust := alpha.Adjust
	ignoreNA := alpha.IgnoreNA
	if alpha.Com != 0 {
		atype = stat.AlphaCom
		param = alpha.Com
	} else if alpha.Span != 0 {
		atype = stat.AlphaSpan
		param = alpha.Span
	} else if alpha.HalfLife != 0 {
		atype = stat.AlphaHalfLife
		param = alpha.HalfLife
	} else {
		atype = stat.AlphaAlpha
		param = alpha.Alpha
	}

	dest := NewSeries(stat.SERIES_TYPE_FLOAT64, this.name, this.Values())
	return stat.ExponentialMovingWindow{
		Data:     dest,
		AType:    atype,
		Param:    param,
		Adjust:   adjust,
		IgnoreNA: ignoreNA,
		Cb:       alpha.Callback,
	}
}
